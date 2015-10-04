
Lambda calculus interpreter
===========================


Data structure
--------------

* define lambda `["lam", arg_name, body]`
* apply lambda `["app", func, arg_value]`
* variable `["var", var_name]`

[
  "app",
  [
    "lam",
    "true",
    [
      "app",
      [
        "lam",
        "false",
        [
          "app",
          [
            "lam",
            "and",
            [
              "app",
              [
                "app",
                [
                  "var",
                  "and"
                ],
                [
                  "var",
                  "true"
                ]
              ],
              [
                "var",
                "true"
              ]
            ]
          ],
          [
            "lam",
            "a",
            [
              "lam",
              "b",
              [
                "app",
                [
                  "app",
                  [
                    "var",
                    "a"
                  ],
                  [
                    "var",
                    "b"
                  ]
                ],
                [
                  "var",
                  "false"
                ]
              ]
            ]
          ]
        ]
      ],
      [
        "lam",
        "a",
        [
          "lam",
          "b",
          [
            "var",
            "b"
          ]
        ]
      ]
    ]
  ],
  [
    "lam",
    "a",
    [
      "lam",
      "b",
      [
        "var",
        "a"
      ]
    ]
  ]
]
