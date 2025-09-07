# 📧 Email Sender Service

Este es un **servicio de envío de emails** implementado como una **AWS Lambda Function** en Go. El proyecto está diseñado para:

## 🎯 Propósito Principal
- Servicio serverless para envío de emails a través de SMTP
- Soporte para emails individuales y envío masivo (batch)
- Integración con AWS Lambda y API Gateway

## 🏗️ Arquitectura
- **Lenguaje**: Go 1.25
- **Framework**: AWS Lambda Go SDK
- **Email Library**: go-mail (wneessen/go-mail)
- **Deployment**: AWS Lambda con API Gateway

## 📂 Estructura del Proyecto
```
├── cmd/main.go              # Entry point de la Lambda
├── config/config.go         # Configuración SMTP desde env vars
├── internal/
│   ├── handler/            # Handlers HTTP para los endpoints
│   ├── models/             # Modelos de datos
│   └── service/            # Lógica de negocio para envío de emails
└── bootstrap               # Binary compilado para AWS Lambda
```

## 🚀 Endpoints Disponibles
1. **`/emailSenderService/send-email`** - Envío de email individual
2. **`/emailSenderService/send-email/batch`** - Envío masivo con worker pool

## ⚙️ Características Técnicas
- **Worker Pool**: Para envío concurrente en batch (5 workers por defecto)
- **Templating**: Soporte para HTML templates con datos dinámicos  
- **Configuración**: Via variables de entorno (SMTP, SSL, timeouts)
- **Logging**: Detallado para debugging
- **Error Handling**: Manejo robusto de errores SMTP

## 🔧 Configuración SMTP
- Compatible con proveedores como Hostinger, Gmail, etc.
- Soporte SSL/TLS configurable
- Autenticación PLAIN
- Timeouts personalizables

## 🔧 Variables de Entorno

```bash
# SMTP Configuration
export SMTP_HOST="your-smtp-host"
export SMTP_PORT="465"
export SMTP_USERNAME="your-email@domain.com"
export SMTP_PASSWORD="your-password"
export SMTP_TIMEOUT="5"
export SMTP_USE_SSL="true"

# API Authentication
export API_TOKEN="your-api-token-here"
```

## 📝 Uso de la API

### Envío Individual
```json
POST /emailSenderService/send-email
{
    "from": "sender@domain.com",
    "to": "recipient@domain.com",
    "subject": "Test Subject",
    "body": "Hello World!",
    "doc_type": "text/plain"
}
```

### Envío Masivo (Batch)
```json
POST /emailSenderService/send-email/batch
{
    "from": "sender@domain.com",
    "subject": "Batch Subject",
    "body": "Hello {{.name}}!",
    "doc_type": "text/html",
    "to_list": [
        {
            "to": "user1@domain.com",
            "template_data": {"name": "User 1"}
        },
        {
            "to": "user2@domain.com", 
            "template_data": {"name": "User 2"}
        }
    ]
}
```

## 🚀 Deployment

1. **Build para AWS Lambda**:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/main.go
   zip bootstrap.zip bootstrap
   ```

2. **Upload a AWS Lambda**:
   - Subir el archivo `bootstrap.zip`
   - Configurar las variables de entorno
   - Configurar API Gateway trigger

Es un microservicio serverless completo para manejo de emails con capacidades de envío masivo optimizado.
