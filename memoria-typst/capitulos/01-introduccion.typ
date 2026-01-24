= Introducción

// resumen del capítulo
#block(
    fill: rgb("#f0f0f0"),
    inset: 1em,
    radius: 4pt,
)[
    *Resumen:* En este capítulo se presenta la motivación del proyecto, los objetivos que se persiguen y la estructura del documento.
]

== Motivación

Texto sobre la motivación del proyecto...

== Objetivos

Los objetivos principales de este Trabajo de Fin de Grado son:

1. Primer objetivo
2. Segundo objetivo
3. Tercer objetivo

== Estructura del documento

Este documento está estructurado de la siguiente manera:

- *Capítulo 2*: Estado de la Cuestión
- *Capítulo 3*: Descripción del Trabajo
- *Capítulo 4*: Conclusiones y Trabajo Futuro

== Ejemplo de figura

Aquí puedes ver cómo insertar una figura:

// #figure(
//   image("../imagenes/arquitectura.png", width: 80%),
//   caption: [Arquitectura del sistema]
// ) <fig:arquitectura>

// Y luego referenciarla: Como se puede ver en la @fig:arquitectura...

== Ejemplo de código

```go
func main() {
    fmt.Println("Hola desde Typst!")
}
```

== Ejemplo de tabla

#figure(
    table(
        columns: 3,
        [*Tecnología*], [*Versión*], [*Propósito*],
        [Go], [1.21+], [Backend],
        [SvelteKit], [2.0], [Frontend],
        [PostgreSQL], [15], [Base de datos],
    ),
    caption: [Tecnologías utilizadas]
)
