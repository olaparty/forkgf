package data

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/7RaeTiU6/9+sq8hSyohYpAxtkiyFLKMfYtSGgxxGGHsLShlSUqh0nZQpGObiCMplX3LlrQoCpF9ZAv9rkjed4wl5/ftuuIfz33fn/t9tvfz3oZoSioOQAfowCDW0RBA/q0D9MDWFWfv6ICa/SWBd3VxNjOlBms6dtofEqvQRaK0xMsMM8wM62pKa/QaUShUE0q7ClmqLa6NNG8Lk9DRKxM/325saGioVSZebQ7uycpKN8m8lHkpkyUtI5fTIGVIgUYghB0SbS7KmVJQAfDjhyGalu76swt25gAAewDA4vKY5+W5+Eo44hz/X5UZzyszn1c2FqimSKoMgNM2ksZQZTS/lc1IGmumOPxzMPRPFi9rw+/BSFtPD7yry//MfKP5Es3mSzRVl3u7vPkcC1T+D56BybzAffMCJWnbTy//DJhJBS77KAC47HYTViQdoAd2ju5SKA9Pm5nhze9SDq38SXJBhv/8j8RjPfBSEngf/LxLEnp1SG2khJSpVnXNNu2qbWt+l1nw4NLO9QAAtiU5GOY4ZrDnccljhlCqTLPPmLOUdVS/MFfvmPR/c0wa5pg0eceMF1TnH8lLuWLHpGcck4Y7thBTT4VQsmLHpFex3JkAPXCU2oFDYXHwFd5WrlVVVrYvw8xwCEhdzi0nhrsbq04eZVWjZ50TZ+uNVOUDAPCvjMEJsxiD37sd4271n0si7vPJSxx7l3so9P6uJyII6q1Eu5KsThP/y2DNHKeRA7+XFABg28o43T0X48yT/GHzPHjzermDr/0ytX8Riud3PjsfzqwTAH7XuCPvLyQAQGRJPpY5Pr8jSDV9OOV8lQTvwvo7/tGzXj6lULyrdOtNAhXdHFVXS3mEEABgy4qpTPf9P1KRn1Y/qVaxEBl/DUXaObqjsLhVrkUIAuoI1tnZdakpOj9H3L862XIAAFj/iMDb1d3ZjoRArAZVrm2UYWbIQAsluMHmiSAlWMYDJ8x/9cAJ88sDX8wCD3KMF66h+XWaby9myT9ztP8J2awfMLIZOzKNDf0Itm+I95Qgy+TmbmtnrpnZuWJH3D1X4QgnHOGXI9CDf5EFPu+G87EtIptn6v0Tolk3YDeMubmRl0G6aWC/9j5av+z8WAulmNk4VmHIhgUgvzxx8nDF/daa0vHzmmM46w3BeH5/cFrze1IXbGHMW36WkOGbtYYc36xFhPkJY5U9bxJDL3qalG95k0z3rcIk7gUgv0zygUxuQmZ6aZWRCBopVl1Z88BMukHG0FirCl1tUFaJzjETQ2rXZxsScj4RymSzSh98+llfRtnM2vOfNdNleMsuiU3HMxhUf5fI6scVaQAAMPhTdbOW/kd1ojOr1X/WerkJVt7YCGULav45cRPCutH6JOIWngF0EHGruF6wzQ23d3TG/u/uGHAaJ8xK98jZY79IZ3dX2XnHmqQRkTUUc8QcxRUWPy8aoksSs8KI3T1XshUtetd4ZoyaEgcAIJZ99ZmnnL1wLFx8C5Y6XjWofVxDQfRw4eb6mwa9Rm0p1L9vmcqGqIcSAADxJYnZSYlN9/2/L6AQXd8TRvph7y/NiMysPTuQ6zxc6De/S5WpltdjAACOS85aeqjUVb5HuLu64lG2Hh6r2G/WQYajPPC+zliJOaCfRtUZJ8/sj1oVNdtqkdX/6GQbS8qJJWfQMMYkml/+Itxhd7fTUnRrUkd0IpelRyPF7zkp1mLTvR0AIL2keIY5dkcXjAN2FfLZYQAoZ1cHV4mjOIff+p1jfXc0/LWpJEk6OlRgZ8zLK+qW9pWnudUwkiMupcay9mWl4vbC/s0l7obszQ84yqK2HzAVOK3AQRD6HmVb7hV15/qou36/b41nC9GqiVjSP31cZXS0IP37dGdo5kNZLirtIAACbtPTuDACwPpo9OMJAPg9lPMoAOjBvNn4lVLn0Z2fZ7fK5k8vqMDa0J4WpYBhXTdw2FV0w9dS9yxJzTZHDpHzqh6IPVMi+Kd9u8e3TIngGS4GU343yk/WTby19jMLdVS5et8ZDOsA37pX2zeHJARvfx1QzB/U+LeToVmIBB0TTQI31bo6LpdhSgdWrZ2sLFzh508XHjt+/O8Ch79li824rdBoxrthAVtYxZgotIO7DLCtk6bJlZdOCo+78speDHW40CUuHcHQMlX8b33aAzz/Bb5rKjWRO0WHn5xLruQ72OrIcYXv/NpTamaYHzRsAyelhr8ypavoCH7EqKMEgrWwUpijIpPWClxX1Se8c96rnHzTNR2pNm70RYuShq1EIN5NPF5iD2fanfiUtPiUwGDBCw5sWrl5MgwjLLVPWmVpYvYS6xrfZlNuurhB3Xrjp43szoeTRuQnz37hy6eOUBUe/3Hr8U7jg5S+5sSzUx/UTz6n99kaWnz7o0WeHUOfGNvepEMahdPrJ+Pcc/wTOU/pJhKOPlKrD044/GP99/4AzI8rA7e+Un7PVMnj5+hJP9hAjXw03E8JVAuFjyj9QztwKYGRWVHEbVzsm1j23eMqOWZWu4/t/yJi+OG8gVTPiKD/o0OCJv4hBVMbvk8jzOWV28I9NTKDNZ0fspzbJX/3xZndApFBCKX+igphujGLK6fufx0fCDyk9OBdFnWnS1rMuYB8oZaeSqehnfK9Z+p9eomjtvvQ1Q5hsbpf/D7jzp7ujLG+Ynn9K2gxHhG0E41lVEhnO3Mi1uJkQdCWhtLYJkuV+5aSZpyb3vnbDIx62lx4IX/l3WU6pbwvxA1itIyXs2U34rTzsm7y8agp1ga4BPP60NUOeLkKYxnTvr44jIsro2hqxvO3I+P3ve42Dm5KfH41hKnilMDVnB4xwVf5m8XT2IrcTl0MOy+tg2D2SeJswOtphiOLNOiuh5qaP37tx7O/T7Y3qI3VS1nz/V4Vcd1U0+mq2vO56jj2QINzGMGU9tp6Hvv7RyptbQPMfY2FakVzqI3bnqgrx3zQlH4+ZR88Uv2Gm00Ubz3McsSwje5xReWjhM8Xf6A36rgHnjQIMx66S8N9cDueS8LrfndnZDCPnmD952gcUXtANOv7t9pXHlG3H8ZMbdDQ+nYuhGXPgy43lZqJAU5uo620rP+WI5CONWrPH1M7dB9N0vU5r03jrPAyx8BboG7c+FF6X5fpTh6Lb2UHMQNWrp1OVqkPRp4PixSuC89jpn30XHiNWq+Kwmmr6afPY68iULa9hJH1B1+bsg+uieMpD7coutqK/WgqPoxvvZePfjkY2C0ldeRzerDaVgdE+GvvjQirl4gj6TI4/j24orhqpwHVQtozOoLMjROb37GMC/xlvxcR3GprIZ81/N6MX8W8M+kWVoH76DUF9bsjsQ1yu7d8s2VMfKjae6bxr9I8xU25CsObKQKfZR5CpzRXqCWsk3TH8BcGnW1i7r/gtumOuZZvtq1io1h6sJtKo2WZ+dnNsTYfjG7u3Jo/2RHxcnrKsLmHdxLVpWyQct/az+FQaE1oz18tFt2EVq1r+DPh3bIjavWGx7W+lSVzvcNPXrzpj2RnP1z2YJ04H59WRZQ/9d6kwfdTTnL+n6yCLZ+YTR8N8y5+i6jw846/YemAOiHuyy7IFJZYk/+Q3sZLYaA4Ze3Hs61Xr01rvi2ZTIj6Xq17ryTksnyUxpn7DOeQD/I+hAw84CshOLrveNB6WLTqQhrH61MEgqDSE//rj/7NfvNtMPjDAXxyAbGtPSCMKvWMhZefeQNvq8ydNTXxkRhV3rGso7ly/s6V397qntCxqgofXK9ZHbk9hDqPmp5vhwBRzF03tle2O2TK+YaDVgtHZE6oJgNv51B2jtbxqiMvHjdx12MH9Fj8iUMK8t47o+W8dGLd1/sGqmVnEffu3btncO8g0eAf1YxJKgPFSsVGuwyNApNYOpYo1xoR8cGWNh2eAj2f3E4n+b+G3na82Vv3dOT7X9h/mZn7/51EFlOarYks5Dp4Tb6JicMkT5ivibdYb/hqo2PKmdJch5MT2c0fr8kpRaMk+W9Lhdhad8bk7dfkTG15iX469gN3jb0uvLT4iF1yeljh9WcjtRvV2N+WMPOZ4+3kPBk8CkwRZsV5N1Is44fCWyKZI+J16Xa5ynGytHS5H+hKe1ohpVi5Gydmb/50jJFtIzUq3gBFredq6reWZvh1vb73EcGLVpqflQQ/qbp4Bx4LEsisai4qajuq+SAp1V6xtK/WKvmF7uV2Z+24qU/x1xKVLD4e7o9x2GJOeN2s0GfFOjAVQsmq6ZySkx8pcElajiW+jw9Jl8PrM5RLEGDf/i3L7t2djvKPzTdDQ5nenOilxKYhLW2UxxwfHn47FfciB7MhtX2/5wcuW8Gwa592pBfF/sP4NZm/aUvRjZ7eWN0nqJyguAgDKrSeuP6rgtemfFua2iVPSx1TDzRhjjDpZHqJodpxsv+szOeiUOXxBB3vKqG91uLja5UenyiVPhA7dG2oYG9GJ1Er8Vq1rNKhvm8yhl5N9uPoi3Vu2RRf8zpaJnLHc/E5DGcJcXt6cxIPlLwPuvPk5I++1sEJHf9nSX+HbWLWRwzeFb3ZlZBbvyPpBnHwXPox76nLvXHtJeUWbZk1dXERdu0V+8zzk2gv9GTJqfXgEbue3zxGzfb99qsKXOSEscCT5J7r6a937Us68rhrs772eAPjs3vq0wFJpxt5DxWcS5Wydabl9LNKPS9Vd+J0oL5lvSndQP8lAd9dYQO6uiwX1eI43nRlxVQr8h1/sxkx2euv0a0RJ8BvwYFXRTcc9OGXl2zd/8L27bo3idzJgx29ee22DaX8OyzM918q9XvTQ5BoMI7mkcphkfbm76jVsvUvfn5PngFJQFSNVkeljX6x387dTfxg5xBagxhn2adfQZ+5fkr1jE/EOl2rOAHlraNf0ZGVpoHB8bHl0m6sIevZA/hLtl0vUL9d+vmG4unvd/FWj5kaDhyg7BFpVDRxP1a9kVfq/bP+7W8ZXLEbKeIKhLmtfdrX4IuFm5rVDQYk3+10KRvDKn8Z1tpKTIgZLbx0Kq91tJKvsuPhk+Zj7UQd3aIMO2EEa9F7r3f+xzJv47OH9yhb04iVDFbzj+Ad23I41/ffFH4kkUaMMBIcVmadbHEfkZpips7PKcS2pmVVlidFDezAtqa1vKq+8s+VU8Y6Vp7xhOEHkoe2Zrndr0B+Q3t+JPbLxERUmOFy8y039RwfczUvG6a9p4coq0Gl+nTcEfH1oGpQSZa+/aK92dpgwsq+2rAtJ10pcxLPLdCcXpOpcPe9LuJS7kvHeBfLKaeQbTdc2ngPpr/CHNRO29qgiTg7UaDB/GJfVyXl83RUecaBEcuIBsUBDaHS/F3u+syjrYEjbwuR/677J+SY0+FH6g7XvGmnYnsPmOAIeRPrWrWPmpz2DrjI87f/bn1cJkeqcDi3L16mdquIcolH/9vgdUVBfkaT/E9aNzUU8764je1zaXrZbPO53NUj6FqCtfvZpFO+iu+VpzKEZAOpdwjkjdT2buuuu65xMZUG2d9sMnbswgsLJGJ98OF3pfUn2Ao27vjwge8GftDGzRPvLHX35ai8Y5aiQSIuKvJD4UR5UIPJDzWB0Om04zS8V0w4fTmln6U9H1Ut6mtpoylqTmWPLt4jNp3GzExhzsKJq/ym9eW1HLcXl6vJCb1Ht7xik33YEv7JyR5pLB5qtEddo+GLduNplfEQw+XRXT8h+QorKbVLqckifUwrO0Cf8dCtzQeY8KonaToaGDd87Erd+0V3Oj/5jF5UZjTGZpyDpXNiyD3m0GOCzf3Qg/20eabhxl5vY827TlXWfRbOPI8f0s/UwmXrEMKDFJU7zzxrqDk15jC6P7xDHc3ryRdNX3P6htnhfslIbCxRXLuaJzwgU4eloCtMNruyOSKcgVVJga2IpkGjHY0l3mpIKrze945HaNCgzuRSYjYrLtTk+iHpbXVCkZ05Q7d3CVp94tmJeiOcd2IIoR87Ush5P0PTgcYqkOL4fVP9o8Ty3EDZqdefej6dfUt1pZhtWGAf5mywDbqAg3oXQ49s37B05AaPHBHWxocjTo/S6S+9d0ejT3TfHIlpLJY+3lCb3QPi3XrUVRjLt9tR+b3GyuhZazyWQXsjy7epnHvZHHKFMVLVg8UtTNk1deRQQJv0Yc24JhU0rfONpM17VJJ4C7Imd1b6pNn272OM7LeJQ6Brzu4dyOlXjtZlQtbemdj3bJr9d9ev/3ZABhMAl5mXa67Pvvng7LA+EkfwkNdb8epal2J+JqqPDtLBRWCQ4oxonupGngyGLc/Mo52tL04bxPQnGmgW7Y8IT3h2ICzcw7R9MrBwu092EoHLnmhIw1fYnzJ292tad+HN8FPyBcwf+BCn+UIL/P069IY2ae+n2ew/UPew0sK/76unPR+9w5PcdAGG+MPBu6tfyPMkyJee8rzR/WTvWN8U41xFmN4x03MAgIfL9vB/VrSK3s3P10A81uWoMwaPJWfIkLgRylykTMustBKtnWoo8uvbqplWRaV2I+Xv5sLjGBEgCADYvORLKyuUzBnj6+r563Pin7268pGBQdm64vAYRxzWHa4/JU1fW6+iEm1mqFNds027shKNNE5JbSeUSWZ5jIwyug0PezATamsksz6lpOlXoe+lzjdMhGjf9ssDAKSWFMNDToy9qyt+CSUVNSjteRnEURa3xTSYFitqLf8ljKyGI1iM3eIayrX0dGc1zPZoJLOIRz3xi+qY8L9Qvx0AgPpzHbO/4ToIRvhB77WhTEfNSkXKqjI/EQKN7oWKelkxMDAwyAldFVK459U3jtMI/fceg3Jo9PgdIbkJVqOMsO7uqIZ3QddfRJlSC10oaBC482a/xKbjIRvktMdFow0TjKUrOKI81oWHu0RHEYk662SIxuGXJZl3lza+2iPbnWQXg7l2MeaWdVAjSKZKwW06HoJ2Hi48QfW7yPLqi1uiAAATfzqJpVcxicl4JU123pDMkfk1t995uIqT5HPMCnnIzQ0CfA7M84grNe9aJQ+5Z5+SVl2JrhbV056bhHUVImVVRsYv76e27yntahe27wjm+Hybftud9s5Y/r/mP5NTjXfS/pyD+ks+Hk5yOlwwjqv5PMm/GNbMDyky/nkOD8tkwe2r1Oud+VC59j9QSS9CRfKkbFR820mpyHf756iWODHWUHBQzmuFxqXWAfrff5cY+PPnEgkuUiBouokZBnRrHogk5zOPQT4HNbcwf+xWWAMWpqLgAqDZow0wAV6/By+ViiLFg0aFOGB4PBRg6fzSUnUxw+p6TApFpjzyHe45jN7daygBuaQRvBpoIogLVo0qZPgiSSNSMGgUiAEGljgHBo8ULWUIFcyQsV8Aq/BBhgqQyw/BpUNzPnAf3CHDF8kPkYJBAz5wH17NgcGDQiv3QYwakKaC4NzQb2tMMG4sNSCfCiKFgMZx4BAlcxCksZ8FKiAJGziEIg0gn+IhhYDGWVhgEPfnIMgEc1aOwkULFs3crPx5BP5CWdG8ZITNy/xfQxfGaeAlQJMv8Kk5DUcgF6chBYOmXOBg1nRguejMymt7CgWDxWTgcqApFrgcCnqwXEyGFAyaV4GDuZKAkYnBrLy2OigYLPAClwPNpHDC5HAwgGUCL6RY0PAJHOs4CdbCTMtSla2FVfYRikUaXiE5NyDBEvipKsQIVhBeIcWDBkfgeNEL8ciEU1ZeJDUTWDx8AhcFjXpww0TpLAAhEz4hhYOGM+BwXxbCLUyLLLUp0cE2pWBmQC7fsfhBwQZT82JuONl8BykONFMBx9m2FiwR4CDFgQYlWOGXBxgOSR6DFAYae4Bfz9azgKUzFqRQ0FgCOwzqAikUSWpiqUdFD3tUiqyAbKhhpZcbf1ZALtQALwQaMYBf5vMgw8mEGpbSwQDTsYMNkM8nkOyJkI4d3FJfGMDCfAIpErRTBj/YPXjBov2+lR/sGnyAtMcGFwBtgcFLCeADS/bYlnKVFeZqLxRpYQMNLgjav+KDCZLlBytuoC14fYJ0pHhgqIXkUMl1NBbsg5DmEhxSYwtYWWOLFBLayoFDVpGDJNenWPlTIQiApTpCcGXQrg1c2TgZmBX5B23QwCExgmBlzR9SSGivBQ5ZRg7yT/3jhPlntBUs27KBy4P2VPhh8m4shkWuZUMKC+2fwGHphMDK2zMrP5g9ILCQjYWa5ucfKANl4M8GAOHn+wr4vwAAAP//lfLuQkM3AAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
