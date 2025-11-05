# Contribuir a Go Koans

¡Gracias por tu interés en contribuir! Este proyecto tiene como objetivo ayudar a desarrolladores Python a aprender Go mediante TDD.

## Cómo Contribuir

### Reportar Bugs o Problemas

1. Verifica que el problema no haya sido reportado antes
2. Crea un issue describiendo:
   - El koan específico
   - El comportamiento esperado vs actual
   - Pasos para reproducir

### Sugerir Mejoras

1. Abre un issue con la etiqueta "enhancement"
2. Describe claramente la mejora propuesta
3. Explica por qué sería útil para el aprendizaje

### Enviar Pull Requests

1. Fork el repositorio
2. Crea una rama desde `main`:
   ```bash
   git checkout -b feature/mi-mejora
   ```

3. Realiza tus cambios siguiendo estos lineamientos:
   - **Nuevos koans**: Deben incluir tests exhaustivos
   - **Comparaciones Python**: Incluye comentarios comparando con Python
   - **Documentación**: Actualiza README.md y GUIA.md si es necesario
   - **Tests**: Asegura que `go test ./...` pase

4. Formatea tu código:
   ```bash
   go fmt ./...
   go vet ./...
   ```

5. Commit con mensajes descriptivos:
   ```bash
   git commit -m "feat: agrega koan sobre sync.Map"
   ```

6. Push y crea el Pull Request

## Estilo de Código

- Usa `gofmt` para formatear
- Sigue las convenciones de Go
- Comentarios en español para mantener consistencia
- TODOs claros y educativos

## Estructura de Koans

Cada koan debe seguir este patrón:

```go
// PASO X: Título descriptivo
// En Python: código equivalente en Python
// En Go: código equivalente en Go

// TODO: Descripción clara de lo que hay que hacer
func FuncionEjemplo() tipo {
    // Implementación con espacios para completar
    return valorPorDefecto
}
```

## Proceso de Revisión

1. Todos los PRs serán revisados
2. Se verificará que:
   - Los tests pasen
   - La documentación esté actualizada
   - El código sea educativo y claro
   - Se mantenga la progresión de dificultad

## Preguntas

Si tienes preguntas, abre un issue con la etiqueta "question".

¡Gracias por ayudar a mejorar Go Koans! 🚀
