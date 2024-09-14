# Notes

    curl --key client.key --cert client.pem -k https://localhost:9090/health

    root@pho3nix:/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/client# curl --key client.key --cert client.pem  https://localhost:9090/health
    curl: (60) SSL certificate problem: self-signed certificate
    More details here: https://curl.se/docs/sslcerts.html

    curl failed to verify the legitimacy of the server and therefore could not
    establish a secure connection to it. To learn more about this situation and
    how to fix it, please visit the web page mentioned above.



En este código, se especifican tres ciphersuites permitidas:

tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305
Explicación de campos importantes:
CipherSuites: Especifica una lista de ciphersuites que el servidor puede utilizar. Puedes encontrar todas las ciphersuites disponibles en la documentación de Go.
PreferServerCipherSuites: Cuando está en true, el servidor tiene la prioridad de elegir la ciphersuite. Si es false, el cliente tiene la preferencia.
Listado de algunas ciphersuites seguras:
tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305
tls.TLS_RSA_WITH_AES_256_GCM_SHA384 (aunque ya no es tan recomendado debido a la falta de Perfect Forward Secrecy)
tls.TLS_AES_128_GCM_SHA256 (para TLS 1.3)
Recomendaciones de seguridad:
Usa TLS 1.2 o superior: TLS 1.2 y 1.3 son las versiones más seguras del protocolo TLS. TLS 1.0 y 1.1 se consideran obsoletas y vulnerables.
Ciphersuites modernas: En la medida de lo posible, utiliza ciphersuites con Perfect Forward Secrecy (PFS), como las basadas en ECDHE (Elliptic Curve Diffie-Hellman).
TLS 1.3: Si tu aplicación puede usar TLS 1.3, esta versión es preferible porque simplifica la lista de ciphersuites y mejora la seguridad y el rendimiento.

MinVersion: tls.VersionTLS13: Esto asegura que la versión mínima de TLS que puede usar el servidor es TLS 1.3.
MaxVersion: tls.VersionTLS13: Esto establece que la versión máxima permitida también es TLS 1.3, forzando que solo se use esta versión.
Ventajas de TLS 1.3:
Mayor seguridad: TLS 1.3 elimina ciphersuites obsoletas y vulnerables.
Mejor rendimiento: TLS 1.3 reduce el número de rondas de handshake, lo que mejora la velocidad de conexión.
Simplificación de ciphersuites: TLS 1.3 tiene un conjunto más limitado y seguro de ciphersuites.
Con esta configuración, tu servidor solo aceptará conexiones que soporten TLS 1.3, garantizando un nivel de seguridad elevado para tu API.


    curl -v --tlsv1.3 https://localhost:9090


    curl -v --tlsv1.3 --key client.key --cert client.pem -k https://localhost:9090/health



    root@pho3nix:/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/client# curl -v --tlsv1.3 --key client.key --cert client.pem -k https://localhost:9090/health
    *   Trying 127.0.0.1:9090...
    * Connected to localhost (127.0.0.1) port 9090 (#0)
    * ALPN, offering h2
    * ALPN, offering http/1.1
    * TLSv1.0 (OUT), TLS header, Certificate Status (22):
    * TLSv1.3 (OUT), TLS handshake, Client hello (1):
    * TLSv1.2 (IN), TLS header, Certificate Status (22):
    * TLSv1.3 (IN), TLS handshake, Server hello (2):
    * TLSv1.2 (IN), TLS header, Finished (20):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.3 (IN), TLS handshake, Certificate (11):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.3 (IN), TLS handshake, CERT verify (15):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.3 (IN), TLS handshake, Finished (20):
    * TLSv1.2 (OUT), TLS header, Finished (20):
    * TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    * TLSv1.3 (OUT), TLS handshake, Finished (20):
    * SSL connection using TLSv1.3 / TLS_AES_128_GCM_SHA256
    * ALPN, server accepted to use h2
    * Server certificate:
    *  subject: C=AR; ST=CABA; L=CABA; O=ExampleOrg; OU=IT Department; CN=*
    *  start date: Mar 17 17:27:09 2024 GMT
    *  expire date: Mar 15 17:27:09 2034 GMT
    *  issuer: C=AR; ST=CABA; L=CABA; O=ExampleOrg; OU=IT Department; CN=*
    *  SSL certificate verify result: self-signed certificate (18), continuing anyway.
    * Using HTTP2, server supports multiplexing
    * Connection state changed (HTTP/2 confirmed)
    * Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    * Using Stream ID: 1 (easy handle 0x56f503681eb0)
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    > GET /health HTTP/2
    > Host: localhost:9090
    > user-agent: curl/7.81.0
    > accept: */*
    > 
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * Connection state changed (MAX_CONCURRENT_STREAMS == 250)!
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    < HTTP/2 200 
    < content-type: text/plain; charset=utf-8
    < content-length: 4
    < date: Sat, 14 Sep 2024 07:30:14 GMT
    < 
    * Connection #0 to host localhost left intact