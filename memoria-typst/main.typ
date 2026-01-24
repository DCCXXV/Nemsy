// TFG - Nemsy
// Marco Antonio Pérez Neira
// Universidad Complutense de Madrid - Facultad de Informática

#import "template.typ": *
#import "config.typ": *

// aplicar plantilla
#show: ucm-tfg.with(
  titulo: titulo,
  titulo-en: titulo-en,
  autor: autor,
  director: director,
  // codirector: codirector,
  convocatoria: convocatoria,
  año: año,
  tipo-documento: tipo-documento,
  grado: grado,
  institucion: institucion,
  resumen: resumen,
  palabras-clave: palabras-clave,
  abstract-en: abstract-en,
  keywords-en: keywords-en,
)

// CAPÍTULOS

// capítulo 1: Introducción
#include "capitulos/01-introduccion.typ"

// capítulo 2: Estado de la Cuestión
#include "capitulos/02-estado-cuestion.typ"

// capítulo 3: Descripción del Trabajo
#include "capitulos/03-descripcion-trabajo.typ"

// capítulo 4: Conclusiones y Trabajo Futuro
#include "capitulos/04-conclusiones.typ"

#pagebreak()
#heading(level: 1, numbering: none)[Bibliografía]

// recordar crear referencias.bib para la bibliografía integrada de typst
// #bibliography("referencias.bib", style: "ieee")

// APENDICES
#pagebreak()
#set heading(numbering: (..nums) => {
    if nums.pos().len() == 1 {
        "Apéndice " + numbering("A", ..nums)
    } else {
        numbering("A.1", ..nums)
    }
})
#counter(heading).update(0)

// Apéndice A
// #include "apendices/a-manual-usuario.typ"

// Apéndice B
// #include "apendices/b-instalacion.typ"
