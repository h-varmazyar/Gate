DROP INDEX idx_candles_market_id;
DROP INDEX idx_candles_resolution_id;

CREATE INDEX idx_candles_market_resolution_time ON public.candles (resolution_id, market_id, time DESC);

DROP INDEX idx_candles_deleted_at;

-- ایجاد جدول اصلی پارتیشن‌شده
CREATE TABLE public.candles_partitioned (LIKE public.candles);

ALTER TABLE public.candles_partitioned DROP CONSTRAINT candles_pkey;
ALTER TABLE public.candles_partitioned ADD PRIMARY KEY (id, time); -- ترکیب id و time به عنوان کلید اصلی

-- ایجاد پارتیشن‌ها (مثال برای پارتیشن‌بندی ماهانه)
CREATE TABLE public.candles_y2023m01 PARTITION OF public.candles_partitioned
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');
CREATE TABLE public.candles_y2023m02 PARTITION OF public.candles_partitioned
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');
-- ... و به همین ترتیب برای ماه‌های دیگر

-- انتقال داده‌ها از جدول اصلی به پارتیشن‌ها
INSERT INTO public.candles_partitioned SELECT * FROM public.candles;

-- حذف جدول اصلی (بعد از اطمینان از صحت انتقال داده‌ها)
DROP TABLE public.candles;

-- تغییر نام جدول پارتیشن‌شده به نام اصلی
ALTER TABLE public.candles_partitioned RENAME TO candles;