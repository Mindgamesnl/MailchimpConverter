# MailchimpConverter
A simple tool that converts your HTML mail templates to mail-ready files (inline CSS, compression etc) while keeping your directory structure. Based on the mailchimp API and tools.

# Usage
simply drop the `mailconverter` binary for your platform in your project, and write your HTML email templates like you would any other web page.
Then simply execute the binary (`./mailconverter`) and it'll load your configured JSON files, inline all your styles, and then write them to the target directory without affecting your source files.

# Configuration
the mailconverter drops a `mail.json` file whenever its first executed, you only need to list your template file, and where you would like the compiled files to be placed, and the tool will do the rest

Example:
```json
[
 {
  "from": "templates/tailwind/welcome.html",
  "to": "templates/mail/welcome.html.twig"
 }
]
```