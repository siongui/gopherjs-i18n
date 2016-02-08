===============
GopherJS_ i18n_
===============

`gettext function`_ in your browser.

This package includes offline code to convert PO_ to JSON_,
and runtime code to translate strings.

Development Environment: `Ubuntu 15.10`_ and `Go 1.5.3`_.


Usage
+++++

Install package:

.. code-block:: bash

  go get -u github.com/siongui/gopherjs-i18n

offline: example for converting PO_ to JSON_:

.. code-block:: go

  package main

  import "github.com/siongui/gopherjs-i18n/tool"

  func main() {
          po2json.PO2JSON("messages", "../pali/common/locale/", "po.json")
  }

*PO2JSON* takes three arguments:

  - *domain*: usually **messages**

  - *localedir*: the directory where you put PO_ files

  - *jsonPath*: output path of JSON_ file

runtime: See `example directory <example/>`_ for demo code.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `[Golang] gettext Function on Frontend (Browser) by GopherJS <https://siongui.github.io/2016/01/28/go-gettext-function-frontend-browser-by-gopherjs/>`_

.. [2] `[Golang] GopherJS Synonyms with JavaScript <https://siongui.github.io/2016/01/29/go-gopherjs-synonyms-with-javascript/>`_

.. [3] `golang list file in directory <https://www.google.com/search?q=golang+list+file+in+directory>`_

       `ReadDir - ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/#ReadDir>`_

.. [4] `golang path join <https://www.google.com/search?q=golang+path+join>`_

       `path - The Go Programming Language <https://golang.org/pkg/path/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _i18n: https://www.google.com/search?q=i18n
.. _gettext function: http://linux.die.net/man/3/gettext
.. _Ubuntu 15.10: http://releases.ubuntu.com/15.10/
.. _Go 1.5.3: https://golang.org/dl/
.. _UNLICENSE: http://unlicense.org/
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _JSON: http://www.json.org/
