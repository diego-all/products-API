## Security features




<details><summary><code> API1:2023 Broken Object Level Authorization </code></summary>

Object level authorization is an access control mechanism that is usually implemented at the code level to validate that a user can only access the objects that they should have permissions to access.

Every API endpoint that receives an ID of an object, and performs any action on the object, should implement object-level authorization checks. The checks should validate that the logged-in user has permissions to perform the requested action on the requested object.

Failures in this mechanism typically lead to unauthorized information disclosure, modification, or destruction of all data.

Comparing the user ID of the current session (e.g. by extracting it from the JWT token) with the vulnerable ID parameter isn't a sufficient solution to solve Broken Object Level Authorization (BOLA). This approach could address only a small subset of cases.

In the case of BOLA, it's by design that the user will have access to the vulnerable API endpoint/function. The violation happens at the object level, by manipulating the ID. If an attacker manages to access an API endpoint/function they should not have access to - this is a case of Broken Function Level Authorization (BFLA) rather than BOLA.



</summary></details>


<details><summary><code> API2:2023 Broken Authentication </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>



<details><summary><code> API3:2023 Broken Object Property Level Authorization </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>



<details><summary><code> API4:2023 Unrestricted Resource Consumption </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>



</summary></details>



<details><summary><code> API5:2023 Broken Function Level Authorization </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>



<details><summary><code> API6:2023 Unrestricted Access to Sensitive Business Flows </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>


<details><summary><code> API7:2023 Server Side Request Forgery </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>



<details><summary><code> API8:2023 Security Misconfiguration </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>


<details><summary><code> API9:2023 Improper Inventory Management </code></summary>


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services |           | &#9744; | &#9744;  |
| The latest security patches are missing, or the systems are out of date                  |          | &#9744; | &#9744;  |
| Unnecessary features are enabled (e.g. HTTP verbs, logging features)                     |   Yes       | &#9745; | &#9744;  |
| There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain |          | &#9744; | &#9744;  |
| Transport Layer Security (TLS) is missing                                                |   Yes      | &#9745; | &#9745;  |
| Security or cache control directives are not sent to clients                             |          | &#9744; | &#9744;  |
| A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set               |    Yes      | &#9745; | &#9744;  |
| Error messages include stack traces, or expose other sensitive information               |          | &#9744; | &#9744;  |


</summary></details>





<details><summary><code> API10:2023 Unsafe Consumption of APIs   </code></summary>


Developers tend to trust data received from third-party APIs more than user input. This is especially true for APIs offered by well-known companies. Because of that, developers tend to adopt weaker security standards, for instance, in regards to input validation and sanitization.

The API might be vulnerable if:


| Item                                                                                     | Relevant | Apply  | Status  |
|------------------------------------------------------------------------------------------|----------|--------|---------|
| Interacts with other APIs over an unencrypted channel                                    |          | &#9744; | &#9744;  |
| Does not properly validate and sanitize data gathered from other APIs prior to processing it or passing it to downstream components |          | &#9744; | &#9744;  |
| Blindly follows redirections                                                              |          | &#9744; | &#9744;  |
| Does not limit the number of resources available to process third-party services responses |          | &#9744; | &#9744;  |
| Does not implement timeouts for interactions with third-party services                   |          | &#9744; | &#9744;  |


</summary></details>


### Other security practices


