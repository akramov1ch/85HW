# Uy vazifasi: Redis va Go yordamida "Rate Limitingni" Amalda Qo'llash

## Maqsad
Ushbu vazifaning maqsadi `Go` web-ilovasida `Redis` yordamida ilg'or rate limiting mexanizmini joriy etish va undan foydalanishni o'rganishdir. Vazifa sizga haqiqiy dunyo amaliyotiga qaratilgan kengaytirilgan rate limiting tizimini yaratishda yordam beradi.


## Talablar
1. Loyiha Sozlamalari:
    - `net/http` dan foydalanib `Go` loyihasini yarating.
    - Ma'lumotlarni saqlash uchun `Redis` sozlang.

2. **Rate Limiting** Middleware:
    - Foydalanuvchi `IP` manziliga asoslanib so'rovlarni kuzatish uchun middleware yarating.
    - Qisqa vaqt oralig'ida (2 daqiqa) 10 tagacha so'rov uchun `burst`ni qabul qiling.
    - So'rovlar soni va vaqtini saqlash uchun `Redis`'dan foydalaning.

3. Dinamik Rate Limiting Sozlamalari:
    - Har bir foydalanuvchi uchun turli `rate limitl`arni sozlashga imkon beradigan dinamik sozlanadigan `API` endpointini yarating.
    - `IP` bo'yicha rate limitlarni sozlash uchun admin endpointini amalga oshiring.
  





