# Vatansoft Stock Backend
Bu case'de stok için ürünün CRUD işlemlerini , filtreleme , kategori ekleme , bütün ürünleri getirme gibi özellikler barındırır. Bunun yanı sıra fatura , ürün özellikleri ve kategoriler içinde CRUD işlemlerini barındırır.

Projede redis (cache) , mongodb(log) , mysql(veritabanı işlemleri) teknolojileri kullanıldı. Proje http://45.12.81.218/ adresinde 1323 portunun üzerinde bir sunucuda çalışmaktadır.
Sunucuya bağlanıp vatansoft klasörü altında dosyalara erişebilirsiniz :

machine host: http://45.12.81.218/

machine user: root

machine pass: mgcHSiXlOKaNd0rx!diyo@

Projeyi kendi bilgisayarınızda çalıştırmak için cmd klasörünün altındaki main.go dosyasını çalıştırabilirsiniz, proje uzak sunucudaki REDİS,MONGODB VE MYSQL'e bağlanacak şekilde ayarlanmıştır.

Proje içinde bu layout kullanılmıştır => https://github.com/golang-standards/project-layout

Proje'de hata yönetimi olarak her katmanda ERRORS paketi kullanılmıştır , belirli bir formatta çıktı sağlanmıştır.

MongoDB'de logdb altında logs koleksiyonunda loglar tutuldu. Loglar id, log mesajı , ip adres ve remote adres bilgilerini içinde tutuyor.

Bütün API endpointlerinin nasıl kullanılacağını görmek ve test etmek için POSTMAN üzerinden bakabilirsiniz :

https://elements.getpostman.com/redirect?entityId=14030852-4128ba3c-e10f-45ef-a5e1-f7394545dd1d&entityType=collection

