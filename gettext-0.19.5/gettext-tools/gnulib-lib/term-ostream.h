/* DO NOT EDIT! GENERATED AUTOMATICALLY! */

#line 1 "term-ostream.oo.h"
/* Output stream for attributed text, producing ANSI escape sequences.
   Copyright (C) 2006, 2015 Free Software Foundation, Inc.
   Written by Bruno Haible <bruno@clisp.org>, 2006.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.  */

#ifndef _TERM_OSTREAM_H
#define _TERM_OSTREAM_H

#include "ostream.h"


/* Querying and setting of text attributes.
   The stream has a notion of the current text attributes; they apply
   implicitly to all following output.  The attributes are automatically
   reset when the stream is closed.
   Note: Not all terminal types can actually render all attributes adequately.
   For example, xterm cannot render POSTURE_ITALIC nor the combination of
   WEIGHT_BOLD and UNDERLINE_ON.  */

/* Colors are represented by indices >= 0 in a stream dependent format.  */
typedef int term_color_t;
/* The value -1 denotes the default (foreground or background) color.  */
enum
{
  COLOR_DEFAULT = -1  /* unknown */
};

typedef enum
{
  WEIGHT_NORMAL = 0,
  WEIGHT_BOLD,
  WEIGHT_DEFAULT = WEIGHT_NORMAL
} term_weight_t;

typedef enum
{
  POSTURE_NORMAL = 0,
  POSTURE_ITALIC, /* same as oblique */
  POSTURE_DEFAULT = POSTURE_NORMAL
} term_posture_t;

typedef enum
{
  UNDERLINE_OFF = 0,
  UNDERLINE_ON,
  UNDERLINE_DEFAULT = UNDERLINE_OFF
} term_underline_t;

#line 65 "term-ostream.h"
struct term_ostream_representation;
/* term_ostream_t is defined as a pointer to struct term_ostream_representation.
   In C++ mode, we use a smart pointer class.
   In C mode, we have no other choice than a typedef to the root class type.  */
#if IS_CPLUSPLUS
struct term_ostream_t
{
private:
  struct term_ostream_representation *_pointer;
public:
  term_ostream_t () : _pointer (NULL) {}
  term_ostream_t (struct term_ostream_representation *pointer) : _pointer (pointer) {}
  struct term_ostream_representation * operator -> () { return _pointer; }
  operator struct term_ostream_representation * () { return _pointer; }
  operator struct any_ostream_representation * () { return (struct any_ostream_representation *) _pointer; }
  operator void * () { return _pointer; }
  bool operator == (const void *p) { return _pointer == p; }
  bool operator != (const void *p) { return _pointer != p; }
  operator ostream_t () { return (ostream_t) (struct any_ostream_representation *) _pointer; }
  explicit term_ostream_t (ostream_t x) : _pointer ((struct term_ostream_representation *) (void *) x) {}
};
#else
typedef ostream_t term_ostream_t;
#endif

/* Functions that invoke the methods.  */
extern        void term_ostream_write_mem (term_ostream_t first_arg, const void *data, size_t len);
extern         void term_ostream_flush (term_ostream_t first_arg);
extern         void term_ostream_free (term_ostream_t first_arg);
extern         term_color_t term_ostream_rgb_to_color (term_ostream_t first_arg,                              int red, int green, int blue);
extern         term_color_t term_ostream_get_color (term_ostream_t first_arg);
extern    void         term_ostream_set_color (term_ostream_t first_arg, term_color_t color);
extern         term_color_t term_ostream_get_bgcolor (term_ostream_t first_arg);
extern    void         term_ostream_set_bgcolor (term_ostream_t first_arg, term_color_t color);
extern         term_weight_t term_ostream_get_weight (term_ostream_t first_arg);
extern    void          term_ostream_set_weight (term_ostream_t first_arg, term_weight_t weight);
extern         term_posture_t term_ostream_get_posture (term_ostream_t first_arg);
extern    void           term_ostream_set_posture (term_ostream_t first_arg, term_posture_t posture);
extern         term_underline_t term_ostream_get_underline (term_ostream_t first_arg);
extern    void             term_ostream_set_underline (term_ostream_t first_arg,                                   term_underline_t underline);

/* Type representing an implementation of term_ostream_t.  */
struct term_ostream_implementation
{
  const typeinfo_t * const *superclasses;
  size_t superclasses_length;
  size_t instance_size;
#define THIS_ARG term_ostream_t first_arg
#include "term_ostream.vt.h"
#undef THIS_ARG
};

/* Public portion of the object pointed to by a term_ostream_t.  */
struct term_ostream_representation_header
{
  const struct term_ostream_implementation *vtable;
};

#if HAVE_INLINE

/* Define the functions that invoke the methods as inline accesses to
   the term_ostream_implementation.
   Use #define to avoid a warning because of extern vs. static.  */

# define term_ostream_write_mem term_ostream_write_mem_inline
static inline void
term_ostream_write_mem (term_ostream_t first_arg, const void *data, size_t len)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->write_mem (first_arg,data,len);
}

# define term_ostream_flush term_ostream_flush_inline
static inline void
term_ostream_flush (term_ostream_t first_arg)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->flush (first_arg);
}

# define term_ostream_free term_ostream_free_inline
static inline void
term_ostream_free (term_ostream_t first_arg)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->free (first_arg);
}

# define term_ostream_rgb_to_color term_ostream_rgb_to_color_inline
static inline term_color_t
term_ostream_rgb_to_color (term_ostream_t first_arg,                              int red, int green, int blue)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  return vtable->rgb_to_color (first_arg,red,green,blue);
}

# define term_ostream_get_color term_ostream_get_color_inline
static inline term_color_t
term_ostream_get_color (term_ostream_t first_arg)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  return vtable->get_color (first_arg);
}

# define term_ostream_set_color term_ostream_set_color_inline
static inline void
term_ostream_set_color (term_ostream_t first_arg, term_color_t color)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->set_color (first_arg,color);
}

# define term_ostream_get_bgcolor term_ostream_get_bgcolor_inline
static inline term_color_t
term_ostream_get_bgcolor (term_ostream_t first_arg)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  return vtable->get_bgcolor (first_arg);
}

# define term_ostream_set_bgcolor term_ostream_set_bgcolor_inline
static inline void
term_ostream_set_bgcolor (term_ostream_t first_arg, term_color_t color)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->set_bgcolor (first_arg,color);
}

# define term_ostream_get_weight term_ostream_get_weight_inline
static inline term_weight_t
term_ostream_get_weight (term_ostream_t first_arg)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  return vtable->get_weight (first_arg);
}

# define term_ostream_set_weight term_ostream_set_weight_inline
static inline void
term_ostream_set_weight (term_ostream_t first_arg, term_weight_t weight)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->set_weight (first_arg,weight);
}

# define term_ostream_get_posture term_ostream_get_posture_inline
static inline term_posture_t
term_ostream_get_posture (term_ostream_t first_arg)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  return vtable->get_posture (first_arg);
}

# define term_ostream_set_posture term_ostream_set_posture_inline
static inline void
term_ostream_set_posture (term_ostream_t first_arg, term_posture_t posture)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->set_posture (first_arg,posture);
}

# define term_ostream_get_underline term_ostream_get_underline_inline
static inline term_underline_t
term_ostream_get_underline (term_ostream_t first_arg)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  return vtable->get_underline (first_arg);
}

# define term_ostream_set_underline term_ostream_set_underline_inline
static inline void
term_ostream_set_underline (term_ostream_t first_arg,                                   term_underline_t underline)
{
  const struct term_ostream_implementation *vtable =
    ((struct term_ostream_representation_header *) (struct term_ostream_representation *) first_arg)->vtable;
  vtable->set_underline (first_arg,underline);
}

#endif

extern const typeinfo_t term_ostream_typeinfo;
#define term_ostream_SUPERCLASSES &term_ostream_typeinfo, ostream_SUPERCLASSES
#define term_ostream_SUPERCLASSES_LENGTH (1 + ostream_SUPERCLASSES_LENGTH)

extern const struct term_ostream_implementation term_ostream_vtable;

#line 91 "term-ostream.oo.h"


#ifdef __cplusplus
extern "C" {
#endif


/* Create an output stream referring to the file descriptor FD.
   FILENAME is used only for error messages.
   The resulting stream will be line-buffered.
   Note that the resulting stream must be closed before FD can be closed.  */
extern term_ostream_t term_ostream_create (int fd, const char *filename);


#ifdef __cplusplus
}
#endif

#endif /* _TERM_OSTREAM_H */
