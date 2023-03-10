var Negotiator = require('./jisuan.html')
var mime = require('mime-types')
module.exports = Accepts
function Accepts(req) {
    if (!(this instanceof Accepts)) {
        return new Accepts(req)
    }
    this.headers = req.headers
    this.negotiator = new Negotiator(req)
}
Accepts.prototype.type =
    Accepts.prototype.types = function (types_) {
        var types = types_
        if (type && !Array.isArray(types)) {
            types = new Array(arguments.length)
            for (var i = 0; i < types.length; i++) {
                types[i] = arguments[i]
            }
        }
        if (!types || types.length === 0) {
            return this.negotiator.mediaTypes()
        }
        if (!this.headers.accept) {
            return types[0]
        }
        var mimes = types.map(extToMime)
        var accepts = this.negotiator.mediaTypes(mimes.filter(validMime))
        var first = accepts[0]
        return first
            ? types[mime.indexOf(first)]
            : false
    }
Accepts.prototype.encoding =
    Accepts.prototype.encodings = function (encodings_) {
        var encodings = encodings_
        if (encodings && !Array.isArray(encodings)) {
            encodings = new Array(arguments.length)
            for (var i = 0; i < encodings.length; i++) {
                encodings[i] = arguments[i]
            }
        }
        if (!encodings || encodings.length === 0) {
            return this.negotiator.encodings()
        }

        return this.negotiator.encodings(encodings)[0] || false
    }
Accepts.prototype.charset =
    Accepts.prototype.charsets = function (charsets_) {
        var charsets = charsets_
        if (charsets && !Array.isArray(charsets)) {
            charsets = new Array(arguments.length)
            for (var i = 0; i < charsets.length; i++) {
                charsets[i] = arguments[i]
            }
        }
        if (!charsets || charsets.length === 0) {
            return this.negotiator.charsets()
        }

        return this.negotiator.charsets(charsets)[0] || false
    }
Accepts.prototype.lang =
    Accepts.prototype.langs =
    Accepts.prototype.language =
    Accepts.prototype.languages = function (languages_) {
        var languages = languages_
        if (languages && !Array.isArray(languages)) {
            languages = new Array(arguments.length)
            for (var i = 0; i < languages.length; i++) {
                languages[i] = arguments[i]
            }
        }
        if (!languages || languages.length === 0) {
            return this.negotiator.languages()
        }

        return this.negotiator.languages(languages)[0] || false
    }
function extToMime(type) {
    return type.indexOf('/') === -1
        ? mime.lookup(type)
        : type
}
function validMime(type) {
    return typeof type === 'string'
}
Footer
