module.exports = {
  plugins: [
    require('stylelint')({
      "extends": "stylelint-config-standard",
      "rules": {
        "selector-pseudo-class-no-unknown": [true, { ignorePseudoClasses: ["global"] }],
        "number-leading-zero": ["never"],
      }
    }),
    require("postcss-reporter")({
      clearMessages: true,
    }),
    require("postcss-nested"),
    require('postcss-assets')({
      loadPaths: ['./build/img', './build/svg'],
    }),
    require('postcss-svgo')(),
  ],
}
