# 📱 WhatsApp CRM SaaS (Laravel) – SOP & Architecture Guide

---

# 1. OVERVIEW SYSTEM

Sistem ini menggunakan:
- WhatsApp Business Cloud API
- Laravel Backend
- Queue System
- Webhook Listener
- Rule Engine + AI Hybrid Chatbot

Tujuan:
- Customer support automation
- CRM WhatsApp
- Notifikasi transaksi
- Chatbot intelligent system

---

# 2. CORE ARCHITECTURE

## Flow System

User WhatsApp  
→ Meta WhatsApp Cloud API  
→ Webhook Laravel  
→ Message Processor  
→ Queue Worker  
→ Business Logic Layer  
→ Response API  
→ WhatsApp Reply

---

# 3. WHATSAPP BUSINESS CLOUD API

## Komponen
- Access Token
- Phone Number ID
- WhatsApp Business Account (WABA)

## Karakteristik
- Hosted by Meta :contentReference[oaicite:0]{index=0}
- REST API
- Real-time webhook event

---

# 4. LARAVEL CLEAN ARCHITECTURE

## Folder Structure


app/
├── Services/
│ └── WhatsAppService.php
│
├── Actions/
│ ├── SendMessageAction.php
│ ├── ProcessIncomingMessageAction.php
│
├── Jobs/
│ └── SendWhatsAppMessageJob.php
│
├── Http/
│ ├── Controllers/
│ │ └── WebhookController.php
│
├── Repositories/
│ └── MessageRepository.php
│
├── Domain/
│ └── Chatbot/
│ ├── RuleEngine.php
│ ├── IntentClassifier.php


---

# 5. WHATSAPP SERVICE LAYER

```php id="svc001"
class WhatsAppService
{
    public function sendText($to, $message)
    {
        return Http::withToken(config('whatsapp.token'))
            ->post($this->endpoint(), [
                'messaging_product' => 'whatsapp',
                'to' => $to,
                'type' => 'text',
                'text' => ['body' => $message]
            ]);
    }

    private function endpoint()
    {
        return 'https://graph.facebook.com/v19.0/' . config('whatsapp.phone_number_id') . '/messages';
    }
}
6. WEBHOOK PROCESSING (ASYNC DESIGN)
Controller
public function handle(Request $request)
{
    ProcessIncomingMessageAction::dispatch($request->all());

    return response('OK', 200);
}
Action Layer
class ProcessIncomingMessageAction
{
    public function handle($data)
    {
        $message = $data['entry'][0]['changes'][0]['value']['messages'][0] ?? null;

        if (!$message) return;

        SendWhatsAppMessageJob::dispatch($message);
    }
}
7. QUEUE SYSTEM (WAJIB PRODUCTION)

Kenapa pakai queue:

Hindari timeout webhook
Scaling tinggi
Reliability
php artisan queue:work
8. RULE ENGINE CHATBOT
Simple Rule Engine
class RuleEngine
{
    public function handle($text)
    {
        return match (strtolower($text)) {
            'halo' => 'Halo 👋, ada yang bisa saya bantu?',
            'harga' => 'Silakan cek katalog kami',
            default => null
        };
    }
}
9. AI HYBRID CHATBOT (ADVANCED)
Flow:

Rule Engine → fallback AI

if ($response = $ruleEngine->handle($text)) {
    return $response;
}

return $this->aiFallback($text);
10. MESSAGE TYPES

Supported:

text
image
document
template message
11. SECURITY LAYER

Checklist:

Validate webhook signature
Verify token
Rate limit endpoint
Store logs safely
12. DATABASE DESIGN
messages table
id
wa_id
phone_number
message
type
direction (in/out)
status
created_at
13. OBSERVABILITY
Log semua webhook
Track message lifecycle
Store API response
14. ERROR HANDLING
Common Issues
Token expired

→ regenerate access token

Webhook retry duplicate

→ implement message_id deduplication

API timeout

→ move to queue

15. SCALABILITY DESIGN
Queue worker scaling
Redis queue recommended
Horizontal scaling webhook
16. SECURITY BEST PRACTICE
Never expose token in logs
Use env encryption
Validate webhook origin
17. DEPLOYMENT ENVIRONMENT
HTTPS wajib
Use supervisor for queue
Separate staging & production WABA
18. FINAL ARCHITECTURE SUMMARY

WhatsApp User
→ Meta Cloud API
→ Laravel Webhook
→ Queue System
→ Rule Engine / AI
→ WhatsApp Response