;;; go-flycheck.el --- goflymake flycheck checker

;;; Commentary:
;; Flycheck checker for the go programming language using goflymake tool
;;
;; Add the following lines to your .emacs:
;;
;;   (add-to-list 'load-path "$GOPATH/src/github.com/ptrv/goflymake")
;;   (require 'go-flycheck)

;;; Code:


(eval-after-load "flycheck"
  '(progn
    (flycheck-declare-checker go-goflycheck
      "A Go syntax and style checker using the go utility.

See URL `https://github.com/ptrv/goflymake'."
      :command '("goflycheck" source-original)
      :error-patterns '(("^\\(?1:.*\\):\\(?2:[0-9]+\\): \\(?4:.*\\)$" error))
      :modes 'go-mode
      :predicate '(not (buffer-modified-p)))

    (add-to-list 'flycheck-checkers 'go-goflycheck)))

(provide 'go-flycheck)

;;; go-flycheck.el ends here
