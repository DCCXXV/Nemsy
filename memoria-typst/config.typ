// Configuración del documento

// información del TFG
#let titulo = "Nemsy, una plataforma web colaborativa para compartir y acceder a recursos académicos"
#let titulo-en = "Nemsy, a collaborative web platform for sharing and accessing academic resources"
#let autor = "Marco Antonio Pérez Neira"
#let director = "Ramón González del Campo Rodríguez Barbero"
#let convocatoria = "Junio" // Febrero/Junio/Septiembre
#let año = datetime.today().year()

// tipo de documento
#let tipo-documento = "Trabajo de Fin de Grado"
#let grado = "Ingeniería de Software"
#let institucion = [
    Grado en #grado\
    Facultad de Informática\
    Universidad Complutense de Madrid
]

// resumen
#let resumen = [
    Este Trabajo de Fin de Grado describe el diseño e implementación de una aplicación web que permite a estudiantes compartir, buscar y descargar apuntes. La plataforma está construida siguiendo una arquitectura cliente-servidor: un backend desarrollado en Go, elegido por su rendimiento y la solidez de su ecosistema para la construcción de APIs, y un frontend en SvelteKit que proporciona una interfaz limpia y responsiva.
]

#let palabras-clave = (
    "Ingeniería de Software",
    "Aplicación Web",
    "Go (Golang)",
    "SvelteKit",
    "PostgreSQL",
    "REST",
    "GRAPHQL",
    "OAuth2",
)

// abstract
#let abstract-en = [
    This Bachelor's Thesis describes the design and implementation of a web application that allows students to share, search and download notes. The platform is built following a client-server architecture: a backend developed in Go, chosen for its performance and the robustness of its ecosystem for building APIs, and a frontend in SvelteKit that provides a clean and responsive interface.
]

#let keywords-en = (
  "Software Engineering",
  "Web Application",
  "Go (Golang)",
  "SvelteKit",
  "PostgreSQL",
  "REST",
  "GRAPHQL",
  "OAuth2",
)
