# Ürün Bilgi Servisi

Bu basit Go programı, ürünlerin varlık durumunu ve stok miktarını kontrol etmek için bir HTTP API sunar.

## Nasıl Çalışır

1. **Kurulum**: Projeyi klonlayın ve Go'nun yüklü olduğundan emin olun.
2. **Bağımlılıklar**: Herhangi bir harici bağımlılık gerektirmez.
3. **Sunucuyu Başlatın**: `go run main.go` komutuyla sunucuyu başlatın.
4. **Ürün Bilgilerini Alın**: Tarayıcıdan veya bir API test aracından `http://localhost:8080/product?id=PRODUCT_ID` adresine bir GET isteği yaparak ürün bilgilerini alın.

## API Dökümantasyonu

### `/product` Endpoint'i

- **Method**: GET
- **Parametreler**:
  - `id`: Ürünün ID'si.
- **Cevap**: JSON formatında ürün bilgilerini içeren bir yanıt döndürür.

Örnek istek:
GET http://localhost:8080/product?id=product123


Örnek cevap:
```json
{
  "id": "product123",
  "exists": true,
  "stock": 10
}


![Açıklama](https://github.com/omerfdev/shoppingCheckAPI/blob/master/goProduct.png)
