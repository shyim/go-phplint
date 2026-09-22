// ragel_subtype=go
//
//line internal/php/scanner.rl:1
package php

import (
	"fmt"

	"github.com/shyim/go-phplint/internal/token"
)

//line internal/php/scanner.go:12
const lexer_start int = 129
const lexer_first_final int = 129
const lexer_error int = 0

const lexer_en_main int = 129
const lexer_en_html int = 132
const lexer_en_php int = 139
const lexer_en_property int = 213
const lexer_en_nowdoc int = 220
const lexer_en_heredoc int = 224
const lexer_en_backqote int = 231
const lexer_en_template_string int = 237
const lexer_en_heredoc_end int = 243
const lexer_en_string_var int = 245
const lexer_en_string_var_index int = 251
const lexer_en_string_var_name int = 262
const lexer_en_halt_compiller_open_parenthesis int = 264
const lexer_en_halt_compiller_close_parenthesis int = 268
const lexer_en_halt_compiller_close_semicolon int = 272
const lexer_en_halt_compiller_end int = 276

//line internal/php/scanner.rl:16

func initLexer(lex *Lexer) {

//line internal/php/scanner.go:38
	{
		lex.cs = lexer_start
		lex.top = 0
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line internal/php/scanner.rl:20
}

func (lex *Lexer) Lex() *token.Token {
	eof := lex.pe
	var tok token.ID

	tkn := lex.tokenPool.Get()

	lblStart := 0
	lblEnd := 0

	_, _ = lblStart, lblEnd

//line internal/php/scanner.go:60
	{
		var _widec int16
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 129:
			goto st129
		case 130:
			goto st130
		case 1:
			goto st1
		case 131:
			goto st131
		case 132:
			goto st132
		case 133:
			goto st133
		case 134:
			goto st134
		case 135:
			goto st135
		case 136:
			goto st136
		case 137:
			goto st137
		case 2:
			goto st2
		case 3:
			goto st3
		case 4:
			goto st4
		case 138:
			goto st138
		case 5:
			goto st5
		case 139:
			goto st139
		case 140:
			goto st140
		case 141:
			goto st141
		case 6:
			goto st6
		case 142:
			goto st142
		case 143:
			goto st143
		case 144:
			goto st144
		case 145:
			goto st145
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 146:
			goto st146
		case 147:
			goto st147
		case 148:
			goto st148
		case 149:
			goto st149
		case 150:
			goto st150
		case 151:
			goto st151
		case 152:
			goto st152
		case 153:
			goto st153
		case 154:
			goto st154
		case 155:
			goto st155
		case 156:
			goto st156
		case 157:
			goto st157
		case 11:
			goto st11
		case 158:
			goto st158
		case 12:
			goto st12
		case 159:
			goto st159
		case 13:
			goto st13
		case 14:
			goto st14
		case 160:
			goto st160
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto st40
		case 41:
			goto st41
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 45:
			goto st45
		case 46:
			goto st46
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 66:
			goto st66
		case 67:
			goto st67
		case 68:
			goto st68
		case 69:
			goto st69
		case 70:
			goto st70
		case 71:
			goto st71
		case 72:
			goto st72
		case 73:
			goto st73
		case 161:
			goto st161
		case 162:
			goto st162
		case 163:
			goto st163
		case 164:
			goto st164
		case 165:
			goto st165
		case 74:
			goto st74
		case 166:
			goto st166
		case 75:
			goto st75
		case 76:
			goto st76
		case 167:
			goto st167
		case 77:
			goto st77
		case 168:
			goto st168
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 169:
			goto st169
		case 170:
			goto st170
		case 171:
			goto st171
		case 81:
			goto st81
		case 82:
			goto st82
		case 172:
			goto st172
		case 83:
			goto st83
		case 173:
			goto st173
		case 84:
			goto st84
		case 174:
			goto st174
		case 175:
			goto st175
		case 176:
			goto st176
		case 85:
			goto st85
		case 86:
			goto st86
		case 87:
			goto st87
		case 88:
			goto st88
		case 177:
			goto st177
		case 178:
			goto st178
		case 89:
			goto st89
		case 179:
			goto st179
		case 180:
			goto st180
		case 90:
			goto st90
		case 91:
			goto st91
		case 92:
			goto st92
		case 93:
			goto st93
		case 181:
			goto st181
		case 94:
			goto st94
		case 95:
			goto st95
		case 96:
			goto st96
		case 97:
			goto st97
		case 182:
			goto st182
		case 183:
			goto st183
		case 184:
			goto st184
		case 185:
			goto st185
		case 186:
			goto st186
		case 187:
			goto st187
		case 98:
			goto st98
		case 188:
			goto st188
		case 189:
			goto st189
		case 99:
			goto st99
		case 190:
			goto st190
		case 191:
			goto st191
		case 100:
			goto st100
		case 192:
			goto st192
		case 193:
			goto st193
		case 101:
			goto st101
		case 102:
			goto st102
		case 194:
			goto st194
		case 195:
			goto st195
		case 196:
			goto st196
		case 197:
			goto st197
		case 198:
			goto st198
		case 199:
			goto st199
		case 200:
			goto st200
		case 201:
			goto st201
		case 202:
			goto st202
		case 103:
			goto st103
		case 203:
			goto st203
		case 204:
			goto st204
		case 205:
			goto st205
		case 206:
			goto st206
		case 207:
			goto st207
		case 208:
			goto st208
		case 104:
			goto st104
		case 105:
			goto st105
		case 106:
			goto st106
		case 107:
			goto st107
		case 108:
			goto st108
		case 109:
			goto st109
		case 209:
			goto st209
		case 210:
			goto st210
		case 110:
			goto st110
		case 211:
			goto st211
		case 212:
			goto st212
		case 213:
			goto st213
		case 214:
			goto st214
		case 215:
			goto st215
		case 111:
			goto st111
		case 216:
			goto st216
		case 217:
			goto st217
		case 218:
			goto st218
		case 112:
			goto st112
		case 219:
			goto st219
		case 220:
			goto st220
		case 0:
			goto st0
		case 221:
			goto st221
		case 222:
			goto st222
		case 223:
			goto st223
		case 224:
			goto st224
		case 225:
			goto st225
		case 113:
			goto st113
		case 226:
			goto st226
		case 227:
			goto st227
		case 228:
			goto st228
		case 229:
			goto st229
		case 230:
			goto st230
		case 231:
			goto st231
		case 114:
			goto st114
		case 115:
			goto st115
		case 232:
			goto st232
		case 233:
			goto st233
		case 234:
			goto st234
		case 235:
			goto st235
		case 236:
			goto st236
		case 237:
			goto st237
		case 116:
			goto st116
		case 117:
			goto st117
		case 238:
			goto st238
		case 239:
			goto st239
		case 240:
			goto st240
		case 241:
			goto st241
		case 242:
			goto st242
		case 243:
			goto st243
		case 244:
			goto st244
		case 245:
			goto st245
		case 246:
			goto st246
		case 247:
			goto st247
		case 248:
			goto st248
		case 118:
			goto st118
		case 249:
			goto st249
		case 119:
			goto st119
		case 120:
			goto st120
		case 250:
			goto st250
		case 251:
			goto st251
		case 252:
			goto st252
		case 253:
			goto st253
		case 254:
			goto st254
		case 255:
			goto st255
		case 256:
			goto st256
		case 257:
			goto st257
		case 121:
			goto st121
		case 122:
			goto st122
		case 258:
			goto st258
		case 123:
			goto st123
		case 259:
			goto st259
		case 124:
			goto st124
		case 260:
			goto st260
		case 261:
			goto st261
		case 262:
			goto st262
		case 263:
			goto st263
		case 125:
			goto st125
		case 264:
			goto st264
		case 265:
			goto st265
		case 266:
			goto st266
		case 126:
			goto st126
		case 267:
			goto st267
		case 268:
			goto st268
		case 269:
			goto st269
		case 270:
			goto st270
		case 127:
			goto st127
		case 271:
			goto st271
		case 272:
			goto st272
		case 273:
			goto st273
		case 274:
			goto st274
		case 128:
			goto st128
		case 275:
			goto st275
		case 276:
			goto st276
		case 277:
			goto st277
		case 278:
			goto st278
		case 279:
			goto st279
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 1:
			goto st_case_1
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 138:
			goto st_case_138
		case 5:
			goto st_case_5
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 6:
			goto st_case_6
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 11:
			goto st_case_11
		case 158:
			goto st_case_158
		case 12:
			goto st_case_12
		case 159:
			goto st_case_159
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 160:
			goto st_case_160
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 74:
			goto st_case_74
		case 166:
			goto st_case_166
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 167:
			goto st_case_167
		case 77:
			goto st_case_77
		case 168:
			goto st_case_168
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 172:
			goto st_case_172
		case 83:
			goto st_case_83
		case 173:
			goto st_case_173
		case 84:
			goto st_case_84
		case 174:
			goto st_case_174
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		case 89:
			goto st_case_89
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 181:
			goto st_case_181
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 98:
			goto st_case_98
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 99:
			goto st_case_99
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
		case 100:
			goto st_case_100
		case 192:
			goto st_case_192
		case 193:
			goto st_case_193
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 201:
			goto st_case_201
		case 202:
			goto st_case_202
		case 103:
			goto st_case_103
		case 203:
			goto st_case_203
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 206:
			goto st_case_206
		case 207:
			goto st_case_207
		case 208:
			goto st_case_208
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 110:
			goto st_case_110
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 213:
			goto st_case_213
		case 214:
			goto st_case_214
		case 215:
			goto st_case_215
		case 111:
			goto st_case_111
		case 216:
			goto st_case_216
		case 217:
			goto st_case_217
		case 218:
			goto st_case_218
		case 112:
			goto st_case_112
		case 219:
			goto st_case_219
		case 220:
			goto st_case_220
		case 0:
			goto st_case_0
		case 221:
			goto st_case_221
		case 222:
			goto st_case_222
		case 223:
			goto st_case_223
		case 224:
			goto st_case_224
		case 225:
			goto st_case_225
		case 113:
			goto st_case_113
		case 226:
			goto st_case_226
		case 227:
			goto st_case_227
		case 228:
			goto st_case_228
		case 229:
			goto st_case_229
		case 230:
			goto st_case_230
		case 231:
			goto st_case_231
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 236:
			goto st_case_236
		case 237:
			goto st_case_237
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 238:
			goto st_case_238
		case 239:
			goto st_case_239
		case 240:
			goto st_case_240
		case 241:
			goto st_case_241
		case 242:
			goto st_case_242
		case 243:
			goto st_case_243
		case 244:
			goto st_case_244
		case 245:
			goto st_case_245
		case 246:
			goto st_case_246
		case 247:
			goto st_case_247
		case 248:
			goto st_case_248
		case 118:
			goto st_case_118
		case 249:
			goto st_case_249
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 250:
			goto st_case_250
		case 251:
			goto st_case_251
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 256:
			goto st_case_256
		case 257:
			goto st_case_257
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 258:
			goto st_case_258
		case 123:
			goto st_case_123
		case 259:
			goto st_case_259
		case 124:
			goto st_case_124
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 262:
			goto st_case_262
		case 263:
			goto st_case_263
		case 125:
			goto st_case_125
		case 264:
			goto st_case_264
		case 265:
			goto st_case_265
		case 266:
			goto st_case_266
		case 126:
			goto st_case_126
		case 267:
			goto st_case_267
		case 268:
			goto st_case_268
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
		case 127:
			goto st_case_127
		case 271:
			goto st_case_271
		case 272:
			goto st_case_272
		case 273:
			goto st_case_273
		case 274:
			goto st_case_274
		case 128:
			goto st_case_128
		case 275:
			goto st_case_275
		case 276:
			goto st_case_276
		case 277:
			goto st_case_277
		case 278:
			goto st_case_278
		case 279:
			goto st_case_279
		}
		goto st_out
	tr0:
		lex.cs = 129
//line internal/php/scanner.rl:130
		(lex.p) = (lex.te) - 1
		{
			lex.cs = 132
			lex.ungetCnt(1)
		}
		goto _again
	tr194:
		lex.cs = 129
//line internal/php/scanner.rl:130
		lex.te = (lex.p) + 1
		{
			lex.cs = 132
			lex.ungetCnt(1)
		}
		goto _again
	tr196:
		lex.cs = 129
//line internal/php/scanner.rl:130
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.cs = 132
			lex.ungetCnt(1)
		}
		goto _again
	tr197:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:127
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st129
	st129:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof129
		}
	st_case_129:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:1248
		if lex.data[(lex.p)] == 35 {
			goto tr195
		}
		goto tr194
	tr195:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st130
	st130:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof130
		}
	st_case_130:
//line internal/php/scanner.go:1263
		if lex.data[(lex.p)] == 33 {
			goto st1
		}
		goto tr196
	tr3:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st1
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line internal/php/scanner.go:1285
		switch lex.data[(lex.p)] {
		case 10:
			goto tr2
		case 13:
			goto tr3
		}
		goto st1
	tr2:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st131
	st131:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof131
		}
	st_case_131:
//line internal/php/scanner.go:1310
		goto tr197
	tr4:
		lex.cs = 132
//line internal/php/scanner.rl:143
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.te)
			lex.cs = 139
		}
		goto _again
	tr7:
		lex.cs = 132
//line internal/php/scanner.rl:147
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
			lex.cs = 139
		}
		goto _again
	tr202:
//line internal/php/scanner.rl:137
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(tkn)
			tok = token.T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 132
				goto _out
			}
		}
		goto st132
	tr204:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:137
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(tkn)
			tok = token.T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 132
				goto _out
			}
		}
		goto st132
	tr210:
		lex.cs = 132
//line internal/php/scanner.rl:143
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.te)
			lex.cs = 139
		}
		goto _again
	tr211:
		lex.cs = 132
//line internal/php/scanner.rl:152
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ECHO
			lex.cs = 139
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr213:
		lex.cs = 132
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:147
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
			lex.cs = 139
		}
		goto _again
	st132:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof132
		}
	st_case_132:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:1400
		switch lex.data[(lex.p)] {
		case 10:
			goto tr199
		case 13:
			goto tr200
		case 60:
			goto st136
		}
		goto st133
	tr200:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st133
	tr205:
//line internal/php/scanner.rl:53

		goto st133
	tr207:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st133
	st133:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof133
		}
	st_case_133:
//line internal/php/scanner.go:1445
		switch lex.data[(lex.p)] {
		case 10:
			goto tr199
		case 13:
			goto tr200
		case 60:
			goto st135
		}
		goto st133
	tr199:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st134
	tr206:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st134
	st134:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof134
		}
	st_case_134:
//line internal/php/scanner.go:1486
		switch lex.data[(lex.p)] {
		case 10:
			goto tr206
		case 13:
			goto tr207
		case 60:
			goto tr208
		}
		goto tr205
	tr208:
//line internal/php/scanner.rl:53

		goto st135
	st135:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof135
		}
	st_case_135:
//line internal/php/scanner.go:1505
		switch lex.data[(lex.p)] {
		case 10:
			goto tr199
		case 13:
			goto tr200
		case 60:
			goto st135
		case 63:
			goto tr202
		}
		goto st133
	st136:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr199
		case 13:
			goto tr200
		case 60:
			goto st135
		case 63:
			goto tr209
		}
		goto st133
	tr209:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st137
	st137:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof137
		}
	st_case_137:
//line internal/php/scanner.go:1543
		switch lex.data[(lex.p)] {
		case 61:
			goto tr211
		case 80:
			goto st2
		case 112:
			goto st2
		}
		goto tr210
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 72:
			goto st3
		case 104:
			goto st3
		}
		goto tr4
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 80:
			goto st4
		case 112:
			goto st4
		}
		goto tr4
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 9:
			goto tr7
		case 10:
			goto tr8
		case 13:
			goto tr9
		case 32:
			goto tr7
		}
		goto tr4
	tr8:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st138
	st138:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof138
		}
	st_case_138:
//line internal/php/scanner.go:1610
		goto tr213
	tr9:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line internal/php/scanner.go:1629
		if lex.data[(lex.p)] == 10 {
			goto tr8
		}
		goto tr4
	tr10:
//line internal/php/scanner.rl:161
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st139
	tr12:
		lex.cs = 139
//line NONE:1
		switch lex.act {
		case 10:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 11:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = lex.integerToken(2, 2)
				{
					(lex.p)++
					goto _out
				}
			}
		case 12:
			{
				(lex.p) = (lex.te) - 1

				base := 10
				if lex.data[lex.ts] == '0' {
					base = 8
				}
				lex.setTokenPosition(tkn)
				tok = lex.integerToken(0, base)
				{
					(lex.p)++
					goto _out
				}
			}
		case 13:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = lex.integerToken(2, 16)
				{
					(lex.p)++
					goto _out
				}
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = lex.integerToken(2, 8)
				{
					(lex.p)++
					goto _out
				}
			}
		case 15:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAME_RELATIVE
				{
					(lex.p)++
					goto _out
				}
			}
		case 16:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAME_QUALIFIED
				{
					(lex.p)++
					goto _out
				}
			}
		case 48:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_SL
				{
					(lex.p)++
					goto _out
				}
			}
		case 74:
			{
				(lex.p) = (lex.te) - 1

				tok = lex.identifierToken()
				lex.setTokenPosition(tkn)
				if tok == token.T_HALT_COMPILER {
					lex.cs = 264
				}
				{
					(lex.p)++
					goto _out
				}
			}
		case 80:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('"'))
				lex.cs = 237
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr15:
//line internal/php/scanner.rl:307
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_CONSTANT_ENCAPSED_STRING
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr19:
//line internal/php/scanner.rl:279
		(lex.p) = (lex.te) - 1
		{
			lex.ungetWhile('&')
			lex.setTokenPosition(tkn)
			tok = token.T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr20:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:279
		lex.te = (lex.p) + 1
		{
			lex.ungetWhile('&')
			lex.setTokenPosition(tkn)
			tok = token.T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr24:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:277
		lex.te = (lex.p) + 1
		{
			lex.ungetWhile('&')
			lex.setTokenPosition(tkn)
			tok = token.T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr26:
//line internal/php/scanner.rl:278
		lex.te = (lex.p) + 1
		{
			lex.ungetWhile('&')
			lex.setTokenPosition(tkn)
			tok = token.T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr27:
//line internal/php/scanner.rl:330
		(lex.p) = (lex.te) - 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st139
	tr31:
//line internal/php/scanner.rl:271
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr47:
//line internal/php/scanner.rl:233
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ARRAY_CAST
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr54:
//line internal/php/scanner.rl:239
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING_CAST
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr58:
//line internal/php/scanner.rl:234
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOL_CAST
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr66:
//line internal/php/scanner.rl:236
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOUBLE_CAST
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr73:
//line internal/php/scanner.rl:237
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INT_CAST
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr82:
//line internal/php/scanner.rl:238
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_CAST
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr86:
//line internal/php/scanner.rl:235
		lex.te = (lex.p) + 1
		{
			if lex.versionAtLeast(8, 0) {
				lex.error(fmt.Sprintf("The (real) cast has been removed, use (float) instead"))
			} else {
				lex.setTokenPosition(tkn)
				tok = token.T_DOUBLE_CAST
			}
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr95:
//line internal/php/scanner.rl:241
		lex.te = (lex.p) + 1
		{
			if lex.versionAtLeast(8, 0) {
				lex.error(fmt.Sprintf("The (unset) cast is no longer supported"))
			} else {
				lex.setTokenPosition(tkn)
				tok = token.T_UNSET_CAST
			}
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr99:
//line internal/php/scanner.rl:240
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VOID_CAST
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr100:
//line internal/php/scanner.rl:282
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ELLIPSIS
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr103:
//line internal/php/scanner.rl:165
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr113:
//line internal/php/scanner.rl:258
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && lex.data[lex.ts] == '/' && lex.data[lex.ts+1] == '*' && lex.data[lex.ts+2] == '*' {
				isDocComment = true
			}

			if isDocComment {
				lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
			} else {
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st139
	tr114:
//line internal/php/scanner.rl:167
		(lex.p) = (lex.te) - 1
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}
			lex.setTokenPosition(tkn)
			tok = lex.integerToken(0, base)
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr128:
		lex.cs = 139
//line internal/php/scanner.rl:163
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 132
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr144:
		lex.cs = 139
//line internal/php/scanner.rl:305
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			lex.cs = 213
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr145:
		lex.cs = 139
//line internal/php/scanner.rl:162
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 132
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr148:
		lex.cs = 139
//line internal/php/scanner.rl:295
		(lex.p) = (lex.te) - 1
		{
			tok = lex.identifierToken()
			lex.setTokenPosition(tkn)
			if tok == token.T_HALT_COMPILER {
				lex.cs = 264
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr161:
//line internal/php/scanner.rl:182
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_YIELD_FROM
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr162:
//line internal/php/scanner.rl:179
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_FULLY_QUALIFIED
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr214:
//line internal/php/scanner.rl:330
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st139
	tr225:
//line internal/php/scanner.rl:271
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr244:
		lex.cs = 139
//line internal/php/scanner.rl:327
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('`'))
			lex.cs = 231
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr245:
//line internal/php/scanner.rl:283
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			if lex.looksLikePropertyHookBlock() {
				tok = token.T_PROPERTY_HOOKS
			} else {
				tok = token.ID(int('{'))
			}
			lex.call(139, 139)
			goto _out
		}
		goto st139
	tr247:
//line internal/php/scanner.rl:293
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('}'))
			lex.ret(1)
			goto _out
		}
		goto st139
	tr248:
//line internal/php/scanner.rl:161
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st139
	tr250:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:161
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st139
	tr254:
//line internal/php/scanner.rl:330
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st139
	tr255:
//line internal/php/scanner.rl:271
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr257:
//line internal/php/scanner.rl:219
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr258:
//line internal/php/scanner.rl:220
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_IDENTICAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr259:
		lex.cs = 139
//line internal/php/scanner.rl:328
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('"'))
			lex.cs = 237
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr260:
//line internal/php/scanner.rl:254
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st139
	tr262:
//line internal/php/scanner.rl:184
		lex.te = (lex.p) + 1
		{
			if lex.versionAtLeast(8, 0) {
				lex.setTokenPosition(tkn)
				tok = token.T_ATTRIBUTE
				{
					(lex.p)++
					lex.cs = 139
					goto _out
				}
			} else {
				// On PHP 7 `#[` starts a `#` line comment, not an attribute.
				end := lex.te
				for end < len(lex.data) && lex.data[end] != '\n' {
					if lex.data[end] == '?' && end+1 < len(lex.data) && lex.data[end+1] == '>' {
						break
					}
					end++
				}
				lex.te = end
				lex.p = end - 1
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st139
	tr265:
//line internal/php/scanner.rl:249
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st139
	tr268:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:249
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st139
	tr272:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:244
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st139
	tr273:
//line internal/php/scanner.rl:244
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st139
	tr275:
//line internal/php/scanner.rl:294
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr276:
//line internal/php/scanner.rl:214
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MOD_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr277:
//line internal/php/scanner.rl:279
		lex.te = (lex.p) + 1
		{
			lex.ungetWhile('&')
			lex.setTokenPosition(tkn)
			tok = token.T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr281:
//line internal/php/scanner.rl:277
		lex.te = (lex.p) + 1
		{
			lex.ungetWhile('&')
			lex.setTokenPosition(tkn)
			tok = token.T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr282:
//line internal/php/scanner.rl:202
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOLEAN_AND
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr284:
//line internal/php/scanner.rl:205
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_AND_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr285:
//line internal/php/scanner.rl:279
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetWhile('&')
			lex.setTokenPosition(tkn)
			tok = token.T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr289:
//line internal/php/scanner.rl:208
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MUL_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr290:
//line internal/php/scanner.rl:227
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_POW
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr291:
//line internal/php/scanner.rl:209
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_POW_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr292:
//line internal/php/scanner.rl:216
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INC
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr293:
//line internal/php/scanner.rl:211
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PLUS_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr294:
//line internal/php/scanner.rl:215
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DEC
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr295:
//line internal/php/scanner.rl:212
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MINUS_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr296:
		lex.cs = 139
//line internal/php/scanner.rl:304
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			lex.cs = 213
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr298:
//line internal/php/scanner.rl:207
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_CONCAT_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr299:
//line internal/php/scanner.rl:165
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr302:
//line internal/php/scanner.rl:210
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DIV_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr303:
//line internal/php/scanner.rl:167
		lex.te = (lex.p)
		(lex.p)--
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}
			lex.setTokenPosition(tkn)
			tok = lex.integerToken(0, base)
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr309:
//line internal/php/scanner.rl:166
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = lex.integerToken(2, 2)
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr310:
//line internal/php/scanner.rl:175
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = lex.integerToken(2, 8)
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr311:
//line internal/php/scanner.rl:174
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = lex.integerToken(2, 16)
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr312:
//line internal/php/scanner.rl:201
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PAAMAYIM_NEKUDOTAYIM
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr313:
		lex.cs = 139
//line internal/php/scanner.rl:163
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 132
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr315:
		lex.cs = 139
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:163
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 132
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr318:
//line internal/php/scanner.rl:219
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr319:
//line internal/php/scanner.rl:228
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr320:
//line internal/php/scanner.rl:223
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SL_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr321:
		lex.cs = 139
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:313
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.heredocLabel = lex.data[lblStart:lblEnd]
			lex.setTokenPosition(tkn)
			tok = token.T_START_HEREDOC

			if lex.isHeredocEnd(lex.p + 1) {
				lex.cs = 243
			} else if lex.data[lblStart-1] == '\'' {
				lex.cs = 220
			} else {
				lex.cs = 224
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr322:
//line internal/php/scanner.rl:226
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_SMALLER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr323:
//line internal/php/scanner.rl:218
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SPACESHIP
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr325:
//line internal/php/scanner.rl:217
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOUBLE_ARROW
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr326:
//line internal/php/scanner.rl:221
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr327:
//line internal/php/scanner.rl:222
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_IDENTICAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr328:
//line internal/php/scanner.rl:225
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_GREATER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr330:
//line internal/php/scanner.rl:229
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SR
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr331:
//line internal/php/scanner.rl:224
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SR_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr335:
		lex.cs = 139
//line internal/php/scanner.rl:162
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 132
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr337:
		lex.cs = 139
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:162
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 132
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr338:
//line internal/php/scanner.rl:230
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_COALESCE
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr339:
//line internal/php/scanner.rl:231
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_COALESCE_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr340:
		lex.cs = 139
//line internal/php/scanner.rl:295
		lex.te = (lex.p)
		(lex.p)--
		{
			tok = lex.identifierToken()
			lex.setTokenPosition(tkn)
			if tok == token.T_HALT_COMPILER {
				lex.cs = 264
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr342:
//line internal/php/scanner.rl:178
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_QUALIFIED
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr353:
//line internal/php/scanner.rl:177
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_RELATIVE
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr358:
//line internal/php/scanner.rl:180
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NS_SEPARATOR
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr359:
//line internal/php/scanner.rl:179
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_FULLY_QUALIFIED
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr361:
//line internal/php/scanner.rl:213
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_XOR_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr362:
//line internal/php/scanner.rl:206
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr363:
//line internal/php/scanner.rl:203
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PIPE
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	tr364:
//line internal/php/scanner.rl:204
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOLEAN_OR
			{
				(lex.p)++
				lex.cs = 139
				goto _out
			}
		}
		goto st139
	st139:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof139
		}
	st_case_139:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:2347
		switch lex.data[(lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr216
		case 32:
			goto tr215
		case 33:
			goto st143
		case 34:
			goto tr218
		case 35:
			goto st146
		case 36:
			goto st151
		case 37:
			goto st153
		case 38:
			goto st154
		case 39:
			goto tr223
		case 40:
			goto tr224
		case 42:
			goto st161
		case 43:
			goto st163
		case 45:
			goto st164
		case 46:
			goto tr229
		case 47:
			goto tr230
		case 48:
			goto tr231
		case 58:
			goto st175
		case 59:
			goto tr233
		case 60:
			goto st179
		case 61:
			goto st183
		case 62:
			goto st185
		case 63:
			goto tr237
		case 64:
			goto tr225
		case 66:
			goto tr239
		case 78:
			goto tr240
		case 89:
			goto tr241
		case 92:
			goto st209
		case 94:
			goto st211
		case 96:
			goto tr244
		case 98:
			goto tr239
		case 110:
			goto tr240
		case 121:
			goto tr241
		case 123:
			goto tr245
		case 124:
			goto st212
		case 125:
			goto tr247
		case 126:
			goto tr225
		case 127:
			goto tr214
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr215
				}
			default:
				goto tr214
			}
		case lex.data[(lex.p)] > 31:
			switch {
			case lex.data[(lex.p)] < 49:
				if 41 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 44 {
					goto tr225
				}
			case lex.data[(lex.p)] > 57:
				if 91 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 93 {
					goto tr225
				}
			default:
				goto tr115
			}
		default:
			goto tr214
		}
		goto tr238
	tr251:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

		goto st140
	tr215:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st140
	st140:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof140
		}
	st_case_140:
//line internal/php/scanner.go:2470
		switch lex.data[(lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr249
		case 32:
			goto tr215
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr215
		}
		goto tr248
	tr11:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st141
	tr252:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st141
	st141:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof141
		}
	st_case_141:
//line internal/php/scanner.go:2520
		switch lex.data[(lex.p)] {
		case 10:
			goto tr252
		case 13:
			goto tr253
		case 32:
			goto tr251
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr251
		}
		goto tr250
	tr249:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st6
	tr253:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line internal/php/scanner.go:2564
		if lex.data[(lex.p)] == 10 {
			goto tr11
		}
		goto tr10
	tr216:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st142
	st142:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof142
		}
	st_case_142:
//line internal/php/scanner.go:2586
		if lex.data[(lex.p)] == 10 {
			goto tr11
		}
		goto tr254
	st143:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof143
		}
	st_case_143:
		if lex.data[(lex.p)] == 61 {
			goto st144
		}
		goto tr255
	st144:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof144
		}
	st_case_144:
		if lex.data[(lex.p)] == 61 {
			goto tr258
		}
		goto tr257
	tr218:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:328
		lex.act = 80
		goto st145
	st145:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof145
		}
	st_case_145:
//line internal/php/scanner.go:2621
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
	tr14:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st7
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line internal/php/scanner.go:2654
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 96:
			goto st7
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st7
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st7
			}
		default:
			goto st7
		}
		goto tr12
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		}
		goto st7
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto tr12
		case 92:
			goto st9
		}
		goto st7
	st146:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof146
		}
	st_case_146:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] > 10:
			if 13 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 13 {
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] >= 10:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 91:
			goto tr262
		case 266:
			goto st147
		case 269:
			goto st147
		case 522:
			goto tr263
		case 525:
			goto tr264
		}
		switch {
		case _widec < 11:
			if _widec <= 9 {
				goto st147
			}
		case _widec > 12:
			if 14 <= _widec {
				goto st147
			}
		default:
			goto st147
		}
		goto tr260
	tr267:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st147
	tr269:
//line internal/php/scanner.rl:53

		goto st147
	tr271:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st147
	st147:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof147
		}
	st_case_147:
//line internal/php/scanner.go:2812
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr266
		case 525:
			goto tr267
		}
		if 512 <= _widec && _widec <= 767 {
			goto st147
		}
		goto tr265
	tr266:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st148
	tr270:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st148
	st148:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof148
		}
	st_case_148:
//line internal/php/scanner.go:2892
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr270
		case 525:
			goto tr271
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr269
		}
		goto tr268
	tr263:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st149
	st149:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof149
		}
	st_case_149:
//line internal/php/scanner.go:2958
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr270
		case 525:
			goto tr271
		}
		switch {
		case _widec > 255:
			if 512 <= _widec && _widec <= 767 {
				goto tr269
			}
		default:
			goto tr273
		}
		goto tr272
	tr264:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st150
	st150:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof150
		}
	st_case_150:
//line internal/php/scanner.go:3029
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr263
		case 525:
			goto tr267
		}
		if 512 <= _widec && _widec <= 767 {
			goto st147
		}
		goto tr265
	st151:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof151
		}
	st_case_151:
		if lex.data[(lex.p)] == 96 {
			goto tr255
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr255
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr255
			}
		default:
			goto tr255
		}
		goto st152
	st152:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof152
		}
	st_case_152:
		if lex.data[(lex.p)] == 96 {
			goto tr275
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr275
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr275
				}
			case lex.data[(lex.p)] >= 91:
				goto tr275
			}
		default:
			goto tr275
		}
		goto st152
	st153:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof153
		}
	st_case_153:
		if lex.data[(lex.p)] == 61 {
			goto tr276
		}
		goto tr255
	st154:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof154
		}
	st_case_154:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr279
		case 13:
			goto tr280
		case 32:
			goto st155
		case 36:
			goto tr281
		case 38:
			goto tr282
		case 46:
			goto tr283
		case 61:
			goto tr284
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st155
		}
		goto tr277
	tr21:
//line internal/php/scanner.rl:53

		goto st155
	st155:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof155
		}
	st_case_155:
//line internal/php/scanner.go:3168
		switch lex.data[(lex.p)] {
		case 10:
			goto tr279
		case 13:
			goto tr280
		case 32:
			goto st155
		case 36:
			goto tr281
		case 46:
			goto tr283
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st155
		}
		goto tr277
	tr279:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st156
	tr22:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st156
	st156:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof156
		}
	st_case_156:
//line internal/php/scanner.go:3216
		switch lex.data[(lex.p)] {
		case 10:
			goto tr22
		case 13:
			goto tr23
		case 32:
			goto tr21
		case 36:
			goto tr24
		case 46:
			goto tr25
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr21
		}
		goto tr20
	tr280:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st157
	tr23:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st157
	st157:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof157
		}
	st_case_157:
//line internal/php/scanner.go:3270
		if lex.data[(lex.p)] == 10 {
			goto tr286
		}
		goto tr285
	tr286:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st11
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
//line internal/php/scanner.go:3292
		switch lex.data[(lex.p)] {
		case 10:
			goto tr22
		case 13:
			goto tr23
		case 32:
			goto tr21
		case 36:
			goto tr24
		case 46:
			goto tr25
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr21
		}
		goto tr20
	tr25:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

		goto st158
	tr283:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st158
	st158:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof158
		}
	st_case_158:
//line internal/php/scanner.go:3326
		if lex.data[(lex.p)] == 46 {
			goto st12
		}
		goto tr285
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		if lex.data[(lex.p)] == 46 {
			goto tr26
		}
		goto tr19
	tr223:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st159
	st159:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof159
		}
	st_case_159:
//line internal/php/scanner.go:3350
		switch lex.data[(lex.p)] {
		case 10:
			goto tr29
		case 13:
			goto tr29
		case 39:
			goto tr15
		case 92:
			goto st14
		}
		goto st13
	tr29:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st13
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
//line internal/php/scanner.go:3379
		switch lex.data[(lex.p)] {
		case 10:
			goto tr29
		case 13:
			goto tr29
		case 39:
			goto tr15
		case 92:
			goto st14
		}
		goto st13
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr29
		case 13:
			goto tr29
		}
		goto st13
	tr224:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st160
	st160:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof160
		}
	st_case_160:
//line internal/php/scanner.go:3413
		switch lex.data[(lex.p)] {
		case 9:
			goto st15
		case 32:
			goto st15
		case 65:
			goto st16
		case 66:
			goto st21
		case 68:
			goto st33
		case 70:
			goto st39
		case 73:
			goto st43
		case 79:
			goto st50
		case 82:
			goto st56
		case 83:
			goto st60
		case 85:
			goto st65
		case 86:
			goto st70
		case 97:
			goto st16
		case 98:
			goto st21
		case 100:
			goto st33
		case 102:
			goto st39
		case 105:
			goto st43
		case 111:
			goto st50
		case 114:
			goto st56
		case 115:
			goto st60
		case 117:
			goto st65
		case 118:
			goto st70
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st15
		}
		goto tr255
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 9:
			goto st15
		case 32:
			goto st15
		case 65:
			goto st16
		case 66:
			goto st21
		case 68:
			goto st33
		case 70:
			goto st39
		case 73:
			goto st43
		case 79:
			goto st50
		case 82:
			goto st56
		case 83:
			goto st60
		case 85:
			goto st65
		case 86:
			goto st70
		case 97:
			goto st16
		case 98:
			goto st21
		case 100:
			goto st33
		case 102:
			goto st39
		case 105:
			goto st43
		case 111:
			goto st50
		case 114:
			goto st56
		case 115:
			goto st60
		case 117:
			goto st65
		case 118:
			goto st70
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st15
		}
		goto tr31
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch lex.data[(lex.p)] {
		case 82:
			goto st17
		case 114:
			goto st17
		}
		goto tr31
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch lex.data[(lex.p)] {
		case 82:
			goto st18
		case 114:
			goto st18
		}
		goto tr31
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 65:
			goto st19
		case 97:
			goto st19
		}
		goto tr31
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 89:
			goto st20
		case 121:
			goto st20
		}
		goto tr31
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 9:
			goto st20
		case 32:
			goto st20
		case 41:
			goto tr47
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st20
		}
		goto tr31
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch lex.data[(lex.p)] {
		case 73:
			goto st22
		case 79:
			goto st27
		case 105:
			goto st22
		case 111:
			goto st27
		}
		goto tr31
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch lex.data[(lex.p)] {
		case 78:
			goto st23
		case 110:
			goto st23
		}
		goto tr31
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 65:
			goto st24
		case 97:
			goto st24
		}
		goto tr31
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 82:
			goto st25
		case 114:
			goto st25
		}
		goto tr31
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 89:
			goto st26
		case 121:
			goto st26
		}
		goto tr31
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 9:
			goto st26
		case 32:
			goto st26
		case 41:
			goto tr54
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st26
		}
		goto tr31
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 79:
			goto st28
		case 111:
			goto st28
		}
		goto tr31
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 76:
			goto st29
		case 108:
			goto st29
		}
		goto tr31
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 9:
			goto st30
		case 32:
			goto st30
		case 41:
			goto tr58
		case 69:
			goto st31
		case 101:
			goto st31
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st30
		}
		goto tr31
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 9:
			goto st30
		case 32:
			goto st30
		case 41:
			goto tr58
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st30
		}
		goto tr31
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 65:
			goto st32
		case 97:
			goto st32
		}
		goto tr31
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 78:
			goto st30
		case 110:
			goto st30
		}
		goto tr31
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 79:
			goto st34
		case 111:
			goto st34
		}
		goto tr31
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 85:
			goto st35
		case 117:
			goto st35
		}
		goto tr31
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 66:
			goto st36
		case 98:
			goto st36
		}
		goto tr31
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 76:
			goto st37
		case 108:
			goto st37
		}
		goto tr31
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch lex.data[(lex.p)] {
		case 69:
			goto st38
		case 101:
			goto st38
		}
		goto tr31
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 9:
			goto st38
		case 32:
			goto st38
		case 41:
			goto tr66
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st38
		}
		goto tr31
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 76:
			goto st40
		case 108:
			goto st40
		}
		goto tr31
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 79:
			goto st41
		case 111:
			goto st41
		}
		goto tr31
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 65:
			goto st42
		case 97:
			goto st42
		}
		goto tr31
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch lex.data[(lex.p)] {
		case 84:
			goto st38
		case 116:
			goto st38
		}
		goto tr31
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch lex.data[(lex.p)] {
		case 78:
			goto st44
		case 110:
			goto st44
		}
		goto tr31
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch lex.data[(lex.p)] {
		case 84:
			goto st45
		case 116:
			goto st45
		}
		goto tr31
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 9:
			goto st46
		case 32:
			goto st46
		case 41:
			goto tr73
		case 69:
			goto st47
		case 101:
			goto st47
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st46
		}
		goto tr31
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch lex.data[(lex.p)] {
		case 9:
			goto st46
		case 32:
			goto st46
		case 41:
			goto tr73
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st46
		}
		goto tr31
	st47:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch lex.data[(lex.p)] {
		case 71:
			goto st48
		case 103:
			goto st48
		}
		goto tr31
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch lex.data[(lex.p)] {
		case 69:
			goto st49
		case 101:
			goto st49
		}
		goto tr31
	st49:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch lex.data[(lex.p)] {
		case 82:
			goto st46
		case 114:
			goto st46
		}
		goto tr31
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch lex.data[(lex.p)] {
		case 66:
			goto st51
		case 98:
			goto st51
		}
		goto tr31
	st51:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch lex.data[(lex.p)] {
		case 74:
			goto st52
		case 106:
			goto st52
		}
		goto tr31
	st52:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch lex.data[(lex.p)] {
		case 69:
			goto st53
		case 101:
			goto st53
		}
		goto tr31
	st53:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch lex.data[(lex.p)] {
		case 67:
			goto st54
		case 99:
			goto st54
		}
		goto tr31
	st54:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch lex.data[(lex.p)] {
		case 84:
			goto st55
		case 116:
			goto st55
		}
		goto tr31
	st55:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch lex.data[(lex.p)] {
		case 9:
			goto st55
		case 32:
			goto st55
		case 41:
			goto tr82
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st55
		}
		goto tr31
	st56:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch lex.data[(lex.p)] {
		case 69:
			goto st57
		case 101:
			goto st57
		}
		goto tr31
	st57:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch lex.data[(lex.p)] {
		case 65:
			goto st58
		case 97:
			goto st58
		}
		goto tr31
	st58:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch lex.data[(lex.p)] {
		case 76:
			goto st59
		case 108:
			goto st59
		}
		goto tr31
	st59:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch lex.data[(lex.p)] {
		case 9:
			goto st59
		case 32:
			goto st59
		case 41:
			goto tr86
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st59
		}
		goto tr31
	st60:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch lex.data[(lex.p)] {
		case 84:
			goto st61
		case 116:
			goto st61
		}
		goto tr31
	st61:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch lex.data[(lex.p)] {
		case 82:
			goto st62
		case 114:
			goto st62
		}
		goto tr31
	st62:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch lex.data[(lex.p)] {
		case 73:
			goto st63
		case 105:
			goto st63
		}
		goto tr31
	st63:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch lex.data[(lex.p)] {
		case 78:
			goto st64
		case 110:
			goto st64
		}
		goto tr31
	st64:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch lex.data[(lex.p)] {
		case 71:
			goto st26
		case 103:
			goto st26
		}
		goto tr31
	st65:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch lex.data[(lex.p)] {
		case 78:
			goto st66
		case 110:
			goto st66
		}
		goto tr31
	st66:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch lex.data[(lex.p)] {
		case 83:
			goto st67
		case 115:
			goto st67
		}
		goto tr31
	st67:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch lex.data[(lex.p)] {
		case 69:
			goto st68
		case 101:
			goto st68
		}
		goto tr31
	st68:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch lex.data[(lex.p)] {
		case 84:
			goto st69
		case 116:
			goto st69
		}
		goto tr31
	st69:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch lex.data[(lex.p)] {
		case 9:
			goto st69
		case 32:
			goto st69
		case 41:
			goto tr95
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st69
		}
		goto tr31
	st70:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch lex.data[(lex.p)] {
		case 79:
			goto st71
		case 111:
			goto st71
		}
		goto tr31
	st71:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch lex.data[(lex.p)] {
		case 73:
			goto st72
		case 105:
			goto st72
		}
		goto tr31
	st72:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch lex.data[(lex.p)] {
		case 68:
			goto st73
		case 100:
			goto st73
		}
		goto tr31
	st73:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch lex.data[(lex.p)] {
		case 9:
			goto st73
		case 32:
			goto st73
		case 41:
			goto tr99
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st73
		}
		goto tr31
	st161:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		switch lex.data[(lex.p)] {
		case 42:
			goto st162
		case 61:
			goto tr289
		}
		goto tr255
	st162:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof162
		}
	st_case_162:
		if lex.data[(lex.p)] == 61 {
			goto tr291
		}
		goto tr290
	st163:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof163
		}
	st_case_163:
		switch lex.data[(lex.p)] {
		case 43:
			goto tr292
		case 61:
			goto tr293
		}
		goto tr255
	st164:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		switch lex.data[(lex.p)] {
		case 45:
			goto tr294
		case 61:
			goto tr295
		case 62:
			goto tr296
		}
		goto tr255
	tr229:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st165
	st165:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof165
		}
	st_case_165:
//line internal/php/scanner.go:4339
		switch lex.data[(lex.p)] {
		case 46:
			goto st74
		case 61:
			goto tr298
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr104
		}
		goto tr255
	st74:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		if lex.data[(lex.p)] == 46 {
			goto tr100
		}
		goto tr31
	tr104:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:165
		lex.act = 10
		goto st166
	st166:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof166
		}
	st_case_166:
//line internal/php/scanner.go:4371
		switch lex.data[(lex.p)] {
		case 69:
			goto st75
		case 95:
			goto st77
		case 101:
			goto st75
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr104
		}
		goto tr299
	st75:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		switch lex.data[(lex.p)] {
		case 43:
			goto st76
		case 45:
			goto st76
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr102
		}
		goto tr12
	st76:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr102
		}
		goto tr12
	tr102:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:165
		lex.act = 10
		goto st167
	st167:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof167
		}
	st_case_167:
//line internal/php/scanner.go:4420
		if lex.data[(lex.p)] == 95 {
			goto st76
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr102
		}
		goto tr299
	st77:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr104
		}
		goto tr103
	tr230:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st168
	st168:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof168
		}
	st_case_168:
//line internal/php/scanner.go:4447
		switch lex.data[(lex.p)] {
		case 42:
			goto st78
		case 47:
			goto st147
		case 61:
			goto tr302
		}
		goto tr255
	tr107:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st78
	tr109:
//line internal/php/scanner.rl:53

		goto st78
	tr111:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st78
	st78:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof78
		}
	st_case_78:
//line internal/php/scanner.go:4492
		switch lex.data[(lex.p)] {
		case 10:
			goto tr106
		case 13:
			goto tr107
		case 42:
			goto st80
		}
		goto st78
	tr106:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st79
	tr110:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st79
	st79:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof79
		}
	st_case_79:
//line internal/php/scanner.go:4533
		switch lex.data[(lex.p)] {
		case 10:
			goto tr110
		case 13:
			goto tr111
		case 42:
			goto tr112
		}
		goto tr109
	tr112:
//line internal/php/scanner.rl:53

		goto st80
	st80:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof80
		}
	st_case_80:
//line internal/php/scanner.go:4552
		switch lex.data[(lex.p)] {
		case 10:
			goto tr106
		case 13:
			goto tr107
		case 42:
			goto st80
		case 47:
			goto tr113
		}
		goto st78
	tr231:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:167
		lex.act = 12
		goto st169
	st169:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof169
		}
	st_case_169:
//line internal/php/scanner.go:4576
		switch lex.data[(lex.p)] {
		case 46:
			goto tr304
		case 66:
			goto st82
		case 69:
			goto st75
		case 79:
			goto st83
		case 88:
			goto st84
		case 95:
			goto st81
		case 98:
			goto st82
		case 101:
			goto st75
		case 111:
			goto st83
		case 120:
			goto st84
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr115
		}
		goto tr303
	tr304:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:165
		lex.act = 10
		goto st170
	st170:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof170
		}
	st_case_170:
//line internal/php/scanner.go:4615
		switch lex.data[(lex.p)] {
		case 69:
			goto st75
		case 101:
			goto st75
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr104
		}
		goto tr299
	tr115:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:167
		lex.act = 12
		goto st171
	st171:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof171
		}
	st_case_171:
//line internal/php/scanner.go:4638
		switch lex.data[(lex.p)] {
		case 46:
			goto tr304
		case 69:
			goto st75
		case 95:
			goto st81
		case 101:
			goto st75
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr115
		}
		goto tr303
	st81:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr115
		}
		goto tr114
	st82:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr116
		}
		goto tr12
	tr116:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:166
		lex.act = 11
		goto st172
	st172:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof172
		}
	st_case_172:
//line internal/php/scanner.go:4683
		if lex.data[(lex.p)] == 95 {
			goto st82
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr116
		}
		goto tr309
	st83:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr117
		}
		goto tr12
	tr117:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:175
		lex.act = 14
		goto st173
	st173:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof173
		}
	st_case_173:
//line internal/php/scanner.go:4712
		if lex.data[(lex.p)] == 95 {
			goto st83
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr117
		}
		goto tr310
	st84:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr118
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr118
			}
		default:
			goto tr118
		}
		goto tr12
	tr118:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:174
		lex.act = 13
		goto st174
	st174:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof174
		}
	st_case_174:
//line internal/php/scanner.go:4750
		if lex.data[(lex.p)] == 95 {
			goto st84
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr118
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr118
			}
		default:
			goto tr118
		}
		goto tr311
	st175:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof175
		}
	st_case_175:
		if lex.data[(lex.p)] == 58 {
			goto tr312
		}
		goto tr255
	tr233:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st176
	st176:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof176
		}
	st_case_176:
//line internal/php/scanner.go:4786
		switch lex.data[(lex.p)] {
		case 10:
			goto tr120
		case 13:
			goto tr121
		case 32:
			goto st85
		case 63:
			goto st88
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st85
		}
		goto tr255
	tr123:
//line internal/php/scanner.rl:53

		goto st85
	st85:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof85
		}
	st_case_85:
//line internal/php/scanner.go:4810
		switch lex.data[(lex.p)] {
		case 10:
			goto tr120
		case 13:
			goto tr121
		case 32:
			goto st85
		case 63:
			goto st88
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st85
		}
		goto tr31
	tr120:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st86
	tr124:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st86
	st86:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof86
		}
	st_case_86:
//line internal/php/scanner.go:4856
		switch lex.data[(lex.p)] {
		case 10:
			goto tr124
		case 13:
			goto tr125
		case 32:
			goto tr123
		case 63:
			goto tr126
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr123
		}
		goto tr31
	tr121:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st87
	tr125:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st87
	st87:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof87
		}
	st_case_87:
//line internal/php/scanner.go:4902
		if lex.data[(lex.p)] == 10 {
			goto tr120
		}
		goto tr31
	tr126:
//line internal/php/scanner.rl:53

		goto st88
	st88:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof88
		}
	st_case_88:
//line internal/php/scanner.go:4916
		if lex.data[(lex.p)] == 62 {
			goto tr127
		}
		goto tr31
	tr127:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st177
	st177:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof177
		}
	st_case_177:
//line internal/php/scanner.go:4931
		switch lex.data[(lex.p)] {
		case 10:
			goto tr129
		case 13:
			goto tr314
		}
		goto tr313
	tr129:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st178
	st178:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof178
		}
	st_case_178:
//line internal/php/scanner.go:4956
		goto tr315
	tr314:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st89
	st89:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof89
		}
	st_case_89:
//line internal/php/scanner.go:4975
		if lex.data[(lex.p)] == 10 {
			goto tr129
		}
		goto tr128
	st179:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof179
		}
	st_case_179:
		switch lex.data[(lex.p)] {
		case 60:
			goto tr316
		case 61:
			goto st182
		case 62:
			goto tr318
		}
		goto tr255
	tr316:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:228
		lex.act = 48
		goto st180
	st180:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof180
		}
	st_case_180:
//line internal/php/scanner.go:5006
		switch lex.data[(lex.p)] {
		case 60:
			goto st90
		case 61:
			goto tr320
		}
		goto tr319
	st90:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch lex.data[(lex.p)] {
		case 9:
			goto st90
		case 32:
			goto st90
		case 34:
			goto st91
		case 39:
			goto st95
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr133
	st91:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr134
	tr134:
//line internal/php/scanner.rl:34
		lblStart = lex.p
		goto st92
	st92:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof92
		}
	st_case_92:
//line internal/php/scanner.go:5074
		switch lex.data[(lex.p)] {
		case 34:
			goto tr135
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr12
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr12
				}
			case lex.data[(lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st92
	tr135:
//line internal/php/scanner.rl:35
		lblEnd = lex.p
		goto st93
	st93:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof93
		}
	st_case_93:
//line internal/php/scanner.go:5108
		switch lex.data[(lex.p)] {
		case 10:
			goto tr137
		case 13:
			goto tr138
		}
		goto tr12
	tr137:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st181
	tr141:
//line internal/php/scanner.rl:35
		lblEnd = lex.p
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st181
	st181:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof181
		}
	st_case_181:
//line internal/php/scanner.go:5147
		goto tr321
	tr138:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st94
	tr142:
//line internal/php/scanner.rl:35
		lblEnd = lex.p
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st94
	st94:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof94
		}
	st_case_94:
//line internal/php/scanner.go:5180
		if lex.data[(lex.p)] == 10 {
			goto tr137
		}
		goto tr12
	st95:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr139
	tr139:
//line internal/php/scanner.rl:34
		lblStart = lex.p
		goto st96
	st96:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof96
		}
	st_case_96:
//line internal/php/scanner.go:5215
		switch lex.data[(lex.p)] {
		case 39:
			goto tr135
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr12
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr12
				}
			case lex.data[(lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st96
	tr133:
//line internal/php/scanner.rl:34
		lblStart = lex.p
		goto st97
	st97:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof97
		}
	st_case_97:
//line internal/php/scanner.go:5249
		switch lex.data[(lex.p)] {
		case 10:
			goto tr141
		case 13:
			goto tr142
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr12
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr12
				}
			case lex.data[(lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st97
	st182:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof182
		}
	st_case_182:
		if lex.data[(lex.p)] == 62 {
			goto tr323
		}
		goto tr322
	st183:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof183
		}
	st_case_183:
		switch lex.data[(lex.p)] {
		case 61:
			goto st184
		case 62:
			goto tr325
		}
		goto tr255
	st184:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof184
		}
	st_case_184:
		if lex.data[(lex.p)] == 61 {
			goto tr327
		}
		goto tr326
	st185:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr328
		case 62:
			goto st186
		}
		goto tr255
	st186:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof186
		}
	st_case_186:
		if lex.data[(lex.p)] == 61 {
			goto tr331
		}
		goto tr330
	tr237:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st187
	st187:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof187
		}
	st_case_187:
//line internal/php/scanner.go:5337
		switch lex.data[(lex.p)] {
		case 45:
			goto st98
		case 62:
			goto tr333
		case 63:
			goto st190
		}
		goto tr255
	st98:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		if lex.data[(lex.p)] == 62 {
			goto tr144
		}
		goto tr31
	tr333:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st188
	st188:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof188
		}
	st_case_188:
//line internal/php/scanner.go:5366
		switch lex.data[(lex.p)] {
		case 10:
			goto tr146
		case 13:
			goto tr336
		}
		goto tr335
	tr146:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st189
	st189:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof189
		}
	st_case_189:
//line internal/php/scanner.go:5391
		goto tr337
	tr336:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st99
	st99:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof99
		}
	st_case_99:
//line internal/php/scanner.go:5410
		if lex.data[(lex.p)] == 10 {
			goto tr146
		}
		goto tr145
	st190:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof190
		}
	st_case_190:
		if lex.data[(lex.p)] == 61 {
			goto tr339
		}
		goto tr338
	tr238:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st191
	st191:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof191
		}
	st_case_191:
//line internal/php/scanner.go:5436
		switch lex.data[(lex.p)] {
		case 92:
			goto st100
		case 96:
			goto tr340
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	st100:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr147
	tr147:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:178
		lex.act = 16
		goto st192
	st192:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof192
		}
	st_case_192:
//line internal/php/scanner.go:5494
		switch lex.data[(lex.p)] {
		case 92:
			goto st100
		case 96:
			goto tr342
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr342
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr342
				}
			case lex.data[(lex.p)] >= 91:
				goto tr342
			}
		default:
			goto tr342
		}
		goto tr147
	tr239:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st193
	st193:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof193
		}
	st_case_193:
//line internal/php/scanner.go:5531
		switch lex.data[(lex.p)] {
		case 34:
			goto st7
		case 60:
			goto st101
		case 92:
			goto st100
		case 96:
			goto tr340
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	st101:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if lex.data[(lex.p)] == 60 {
			goto st102
		}
		goto tr148
	st102:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		if lex.data[(lex.p)] == 60 {
			goto st90
		}
		goto tr148
	tr240:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st194
	st194:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof194
		}
	st_case_194:
//line internal/php/scanner.go:5590
		switch lex.data[(lex.p)] {
		case 65:
			goto tr344
		case 92:
			goto st100
		case 96:
			goto tr340
		case 97:
			goto tr344
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr344:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st195
	st195:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof195
		}
	st_case_195:
//line internal/php/scanner.go:5631
		switch lex.data[(lex.p)] {
		case 77:
			goto tr345
		case 92:
			goto st100
		case 96:
			goto tr340
		case 109:
			goto tr345
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr345:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st196
	st196:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof196
		}
	st_case_196:
//line internal/php/scanner.go:5672
		switch lex.data[(lex.p)] {
		case 69:
			goto tr346
		case 92:
			goto st100
		case 96:
			goto tr340
		case 101:
			goto tr346
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr346:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st197
	st197:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof197
		}
	st_case_197:
//line internal/php/scanner.go:5713
		switch lex.data[(lex.p)] {
		case 83:
			goto tr347
		case 92:
			goto st100
		case 96:
			goto tr340
		case 115:
			goto tr347
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr347:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st198
	st198:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof198
		}
	st_case_198:
//line internal/php/scanner.go:5754
		switch lex.data[(lex.p)] {
		case 80:
			goto tr348
		case 92:
			goto st100
		case 96:
			goto tr340
		case 112:
			goto tr348
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr348:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st199
	st199:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof199
		}
	st_case_199:
//line internal/php/scanner.go:5795
		switch lex.data[(lex.p)] {
		case 65:
			goto tr349
		case 92:
			goto st100
		case 96:
			goto tr340
		case 97:
			goto tr349
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr349:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st200
	st200:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof200
		}
	st_case_200:
//line internal/php/scanner.go:5836
		switch lex.data[(lex.p)] {
		case 67:
			goto tr350
		case 92:
			goto st100
		case 96:
			goto tr340
		case 99:
			goto tr350
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr350:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st201
	st201:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof201
		}
	st_case_201:
//line internal/php/scanner.go:5877
		switch lex.data[(lex.p)] {
		case 69:
			goto tr351
		case 92:
			goto st100
		case 96:
			goto tr340
		case 101:
			goto tr351
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr351:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st202
	st202:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof202
		}
	st_case_202:
//line internal/php/scanner.go:5918
		switch lex.data[(lex.p)] {
		case 92:
			goto st103
		case 96:
			goto tr340
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	st103:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr150
	tr150:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:177
		lex.act = 15
		goto st203
	st203:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof203
		}
	st_case_203:
//line internal/php/scanner.go:5976
		switch lex.data[(lex.p)] {
		case 92:
			goto st103
		case 96:
			goto tr353
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr353
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr353
				}
			case lex.data[(lex.p)] >= 91:
				goto tr353
			}
		default:
			goto tr353
		}
		goto tr150
	tr241:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st204
	st204:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof204
		}
	st_case_204:
//line internal/php/scanner.go:6013
		switch lex.data[(lex.p)] {
		case 73:
			goto tr354
		case 92:
			goto st100
		case 96:
			goto tr340
		case 105:
			goto tr354
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr354:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st205
	st205:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof205
		}
	st_case_205:
//line internal/php/scanner.go:6054
		switch lex.data[(lex.p)] {
		case 69:
			goto tr355
		case 92:
			goto st100
		case 96:
			goto tr340
		case 101:
			goto tr355
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr355:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st206
	st206:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof206
		}
	st_case_206:
//line internal/php/scanner.go:6095
		switch lex.data[(lex.p)] {
		case 76:
			goto tr356
		case 92:
			goto st100
		case 96:
			goto tr340
		case 108:
			goto tr356
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr356:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st207
	st207:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof207
		}
	st_case_207:
//line internal/php/scanner.go:6136
		switch lex.data[(lex.p)] {
		case 68:
			goto tr357
		case 92:
			goto st100
		case 96:
			goto tr340
		case 100:
			goto tr357
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr340
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			case lex.data[(lex.p)] >= 91:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr357:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:295
		lex.act = 74
		goto st208
	st208:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof208
		}
	st_case_208:
//line internal/php/scanner.go:6177
		switch lex.data[(lex.p)] {
		case 10:
			goto tr152
		case 13:
			goto tr153
		case 32:
			goto st104
		case 92:
			goto st100
		case 96:
			goto tr340
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st104
				}
			default:
				goto tr340
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 91:
				if 58 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 64 {
					goto tr340
				}
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr340
				}
			default:
				goto tr340
			}
		default:
			goto tr340
		}
		goto tr238
	tr155:
//line internal/php/scanner.rl:53

		goto st104
	st104:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof104
		}
	st_case_104:
//line internal/php/scanner.go:6226
		switch lex.data[(lex.p)] {
		case 10:
			goto tr152
		case 13:
			goto tr153
		case 32:
			goto st104
		case 70:
			goto st107
		case 102:
			goto st107
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st104
		}
		goto tr148
	tr152:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st105
	tr156:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st105
	st105:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof105
		}
	st_case_105:
//line internal/php/scanner.go:6274
		switch lex.data[(lex.p)] {
		case 10:
			goto tr156
		case 13:
			goto tr157
		case 32:
			goto tr155
		case 70:
			goto tr158
		case 102:
			goto tr158
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr155
		}
		goto tr148
	tr153:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st106
	tr157:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st106
	st106:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof106
		}
	st_case_106:
//line internal/php/scanner.go:6322
		if lex.data[(lex.p)] == 10 {
			goto tr152
		}
		goto tr148
	tr158:
//line internal/php/scanner.rl:53

		goto st107
	st107:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof107
		}
	st_case_107:
//line internal/php/scanner.go:6336
		switch lex.data[(lex.p)] {
		case 82:
			goto st108
		case 114:
			goto st108
		}
		goto tr148
	st108:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		switch lex.data[(lex.p)] {
		case 79:
			goto st109
		case 111:
			goto st109
		}
		goto tr148
	st109:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr161
		case 109:
			goto tr161
		}
		goto tr148
	st209:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof209
		}
	st_case_209:
		if lex.data[(lex.p)] == 96 {
			goto tr358
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr358
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr358
			}
		default:
			goto tr358
		}
		goto tr163
	tr163:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st210
	st210:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof210
		}
	st_case_210:
//line internal/php/scanner.go:6399
		switch lex.data[(lex.p)] {
		case 92:
			goto st110
		case 96:
			goto tr359
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr359
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr359
				}
			case lex.data[(lex.p)] >= 91:
				goto tr359
			}
		default:
			goto tr359
		}
		goto tr163
	st110:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		if lex.data[(lex.p)] == 96 {
			goto tr162
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr162
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr162
			}
		default:
			goto tr162
		}
		goto tr163
	st211:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof211
		}
	st_case_211:
		if lex.data[(lex.p)] == 61 {
			goto tr361
		}
		goto tr255
	st212:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof212
		}
	st_case_212:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr362
		case 62:
			goto tr363
		case 124:
			goto tr364
		}
		goto tr255
	tr164:
//line internal/php/scanner.rl:337
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st213
	tr166:
//line internal/php/scanner.rl:341
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				goto st139
			}
		}
		goto st213
	tr167:
//line internal/php/scanner.rl:339
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 213
				goto _out
			}
		}
		goto st213
	tr365:
//line internal/php/scanner.rl:341
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				goto st139
			}
		}
		goto st213
	tr371:
//line internal/php/scanner.rl:337
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st213
	tr373:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:337
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st213
	tr377:
//line internal/php/scanner.rl:341
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				goto st139
			}
		}
		goto st213
	tr378:
//line internal/php/scanner.rl:338
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 213
				goto _out
			}
		}
		goto st213
	tr380:
		lex.cs = 213
//line internal/php/scanner.rl:340
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			lex.cs = 139
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st213:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof213
		}
	st_case_213:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:6531
		switch lex.data[(lex.p)] {
		case 10:
			goto tr165
		case 13:
			goto tr367
		case 32:
			goto tr366
		case 45:
			goto st217
		case 63:
			goto tr369
		case 96:
			goto tr365
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr366
				}
			default:
				goto tr365
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr365
				}
			case lex.data[(lex.p)] >= 91:
				goto tr365
			}
		default:
			goto tr365
		}
		goto st219
	tr374:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

		goto st214
	tr366:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st214
	st214:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof214
		}
	st_case_214:
//line internal/php/scanner.go:6586
		switch lex.data[(lex.p)] {
		case 10:
			goto tr165
		case 13:
			goto tr372
		case 32:
			goto tr366
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr366
		}
		goto tr371
	tr165:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st215
	tr375:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st215
	st215:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof215
		}
	st_case_215:
//line internal/php/scanner.go:6636
		switch lex.data[(lex.p)] {
		case 10:
			goto tr375
		case 13:
			goto tr376
		case 32:
			goto tr374
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr374
		}
		goto tr373
	tr372:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st111
	tr376:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st111
	st111:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof111
		}
	st_case_111:
//line internal/php/scanner.go:6680
		if lex.data[(lex.p)] == 10 {
			goto tr165
		}
		goto tr164
	tr367:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st216
	st216:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof216
		}
	st_case_216:
//line internal/php/scanner.go:6702
		if lex.data[(lex.p)] == 10 {
			goto tr165
		}
		goto tr377
	st217:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof217
		}
	st_case_217:
		if lex.data[(lex.p)] == 62 {
			goto tr378
		}
		goto tr377
	tr369:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st218
	st218:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof218
		}
	st_case_218:
//line internal/php/scanner.go:6726
		if lex.data[(lex.p)] == 45 {
			goto st112
		}
		goto tr377
	st112:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		if lex.data[(lex.p)] == 62 {
			goto tr167
		}
		goto tr166
	st219:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof219
		}
	st_case_219:
		if lex.data[(lex.p)] == 96 {
			goto tr380
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr380
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr380
				}
			case lex.data[(lex.p)] >= 91:
				goto tr380
			}
		default:
			goto tr380
		}
		goto st219
	tr384:
		lex.cs = 220
//line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 87:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				lex.cs = 243
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr385:
		lex.cs = 220
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:345
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.cs = 243
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr389:
		lex.cs = 220
//line internal/php/scanner.rl:345
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.cs = 243
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st220:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof220
		}
	st_case_220:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:6823
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr382
		case 1037:
			goto tr383
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr381
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	tr381:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:345
		lex.act = 87
		goto st221
	tr386:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:345
		lex.act = 87
		goto st221
	st221:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof221
		}
	st_case_221:
//line internal/php/scanner.go:6897
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr382
		case 1037:
			goto tr383
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr381
		}
		goto tr384
	tr382:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st222
	tr387:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st222
	st222:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof222
		}
	st_case_222:
//line internal/php/scanner.go:6977
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr387
		case 1037:
			goto tr388
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr386
		}
		goto tr385
	tr383:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st223
	tr388:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st223
	st223:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof223
		}
	st_case_223:
//line internal/php/scanner.go:7057
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr382
		case 1037:
			goto tr383
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr381
		}
		goto tr389
	tr168:
//line internal/php/scanner.rl:354
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(224, 139)
			goto _out
		}
		goto st224
	tr397:
//line internal/php/scanner.rl:356
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 224
					lex.top++
					goto st245
				}
			}
		}
		goto st224
	tr398:
//line internal/php/scanner.rl:355
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(224, 262)
			goto _out
		}
		goto st224
	tr399:
		lex.cs = 224
//line NONE:1
		switch lex.act {
		case 88:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(224, 139)
				goto _out
			}
		case 89:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(224, 262)
				goto _out
			}
		case 91:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE

				if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
					lex.cs = 243
				}
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr400:
		lex.cs = 224
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:357
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 243
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr404:
		lex.cs = 224
//line internal/php/scanner.rl:357
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 243
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st224:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof224
		}
	st_case_224:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:7189
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto st225
		case 1403:
			goto st113
		case 1546:
			goto tr393
		case 1549:
			goto tr394
		case 1572:
			goto st229
		case 1659:
			goto st230
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr392
		}
		goto st0
	st225:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof225
		}
	st_case_225:
		if lex.data[(lex.p)] == 123 {
			goto tr398
		}
		goto tr397
	st113:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		if lex.data[(lex.p)] == 36 {
			goto tr168
		}
		goto st0
	tr392:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:357
		lex.act = 91
		goto st226
	tr401:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:357
		lex.act = 91
		goto st226
	tr405:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:355
		lex.act = 89
		goto st226
	tr406:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:354
		lex.act = 88
		goto st226
	st226:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof226
		}
	st_case_226:
//line internal/php/scanner.go:7299
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr393
		case 1549:
			goto tr394
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr392
		}
		goto tr399
	tr393:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st227
	tr402:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st227
	st227:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof227
		}
	st_case_227:
//line internal/php/scanner.go:7379
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr402
		case 1549:
			goto tr403
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr401
		}
		goto tr400
	tr394:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st228
	tr403:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st228
	st228:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof228
		}
	st_case_228:
//line internal/php/scanner.go:7459
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr393
		case 1549:
			goto tr394
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr392
		}
		goto tr404
	st229:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof229
		}
	st_case_229:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1403:
			goto tr398
		case 1546:
			goto tr393
		case 1549:
			goto tr394
		case 1659:
			goto tr405
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr392
		}
		goto tr397
	st230:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof230
		}
	st_case_230:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto tr168
		case 1546:
			goto tr393
		case 1549:
			goto tr394
		case 1572:
			goto tr406
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr392
		}
		goto tr404
	tr170:
//line internal/php/scanner.rl:371
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(2)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 231
					lex.top++
					goto st245
				}
			}
		}
		goto st231
	tr171:
//line internal/php/scanner.rl:370
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(231, 262)
			goto _out
		}
		goto st231
	tr172:
//line internal/php/scanner.rl:369
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(231, 139)
			goto _out
		}
		goto st231
	tr408:
		lex.cs = 231
//line internal/php/scanner.rl:372
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('`'))
			lex.cs = 139
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr416:
		lex.cs = 231
//line NONE:1
		switch lex.act {
		case 92:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(231, 139)
				goto _out
			}
		case 93:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(231, 262)
				goto _out
			}
		case 94:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(2)
				{
					lex.growCallStack()
					{
						lex.stack[lex.top] = 231
						lex.top++
						goto st245
					}
				}
			}
		case 95:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('`'))
				lex.cs = 139
				{
					(lex.p)++
					goto _out
				}
			}
		case 96:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr417:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:373
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 231
				goto _out
			}
		}
		goto st231
	tr421:
//line internal/php/scanner.rl:373
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 231
				goto _out
			}
		}
		goto st231
	st231:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof231
		}
	st_case_231:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:7702
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto st114
		case 1888:
			goto tr408
		case 1915:
			goto st115
		case 2058:
			goto tr411
		case 2061:
			goto tr412
		case 2084:
			goto st235
		case 2144:
			goto tr414
		case 2171:
			goto st236
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr410
		}
		goto st0
	st114:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof114
		}
	st_case_114:
		switch lex.data[(lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr171
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr170
	st115:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		if lex.data[(lex.p)] == 36 {
			goto tr172
		}
		goto st0
	tr410:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:373
		lex.act = 96
		goto st232
	tr414:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:372
		lex.act = 95
		goto st232
	tr418:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:373
		lex.act = 96
		goto st232
	tr422:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:371
		lex.act = 94
		goto st232
	tr423:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:370
		lex.act = 93
		goto st232
	tr424:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:369
		lex.act = 92
		goto st232
	st232:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof232
		}
	st_case_232:
//line internal/php/scanner.go:7845
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr411
		case 2061:
			goto tr412
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr410
		}
		goto tr416
	tr411:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st233
	tr419:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st233
	st233:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof233
		}
	st_case_233:
//line internal/php/scanner.go:7925
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr419
		case 2061:
			goto tr420
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr418
		}
		goto tr417
	tr412:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st234
	tr420:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st234
	st234:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof234
		}
	st_case_234:
//line internal/php/scanner.go:8005
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr411
		case 2061:
			goto tr412
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr410
		}
		goto tr421
	st235:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof235
		}
	st_case_235:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1887:
			goto tr170
		case 1915:
			goto tr171
		case 2058:
			goto tr411
		case 2061:
			goto tr412
		case 2143:
			goto tr422
		case 2171:
			goto tr423
		}
		switch {
		case _widec < 2113:
			switch {
			case _widec < 1889:
				if 1857 <= _widec && _widec <= 1882 {
					goto tr170
				}
			case _widec > 1914:
				switch {
				case _widec > 2047:
					if 2048 <= _widec && _widec <= 2112 {
						goto tr410
					}
				case _widec >= 1920:
					goto tr170
				}
			default:
				goto tr170
			}
		case _widec > 2138:
			switch {
			case _widec < 2145:
				if 2139 <= _widec && _widec <= 2144 {
					goto tr410
				}
			case _widec > 2170:
				switch {
				case _widec > 2175:
					if 2176 <= _widec && _widec <= 2303 {
						goto tr422
					}
				case _widec >= 2172:
					goto tr410
				}
			default:
				goto tr422
			}
		default:
			goto tr422
		}
		goto tr421
	st236:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof236
		}
	st_case_236:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto tr172
		case 2058:
			goto tr411
		case 2061:
			goto tr412
		case 2084:
			goto tr424
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr410
		}
		goto tr421
	tr173:
//line internal/php/scanner.rl:383
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(2)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 237
					lex.top++
					goto st245
				}
			}
		}
		goto st237
	tr174:
//line internal/php/scanner.rl:382
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(237, 262)
			goto _out
		}
		goto st237
	tr175:
//line internal/php/scanner.rl:381
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(237, 139)
			goto _out
		}
		goto st237
	tr425:
		lex.cs = 237
//line internal/php/scanner.rl:384
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('"'))
			lex.cs = 139
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr434:
		lex.cs = 237
//line NONE:1
		switch lex.act {
		case 97:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(237, 139)
				goto _out
			}
		case 98:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(237, 262)
				goto _out
			}
		case 99:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(2)
				{
					lex.growCallStack()
					{
						lex.stack[lex.top] = 237
						lex.top++
						goto st245
					}
				}
			}
		case 100:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('"'))
				lex.cs = 139
				{
					(lex.p)++
					goto _out
				}
			}
		case 101:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr435:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:385
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 237
				goto _out
			}
		}
		goto st237
	tr439:
//line internal/php/scanner.rl:385
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 237
				goto _out
			}
		}
		goto st237
	st237:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof237
		}
	st_case_237:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:8289
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2338:
			goto tr425
		case 2340:
			goto st116
		case 2427:
			goto st117
		case 2570:
			goto tr429
		case 2573:
			goto tr430
		case 2594:
			goto tr431
		case 2596:
			goto st241
		case 2683:
			goto st242
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr428
		}
		goto st0
	st116:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof116
		}
	st_case_116:
		switch lex.data[(lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr174
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr173
	st117:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		if lex.data[(lex.p)] == 36 {
			goto tr175
		}
		goto st0
	tr428:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:385
		lex.act = 101
		goto st238
	tr431:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:384
		lex.act = 100
		goto st238
	tr436:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:385
		lex.act = 101
		goto st238
	tr440:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:383
		lex.act = 99
		goto st238
	tr441:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:382
		lex.act = 98
		goto st238
	tr442:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:381
		lex.act = 97
		goto st238
	st238:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof238
		}
	st_case_238:
//line internal/php/scanner.go:8432
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr429
		case 2573:
			goto tr430
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr428
		}
		goto tr434
	tr429:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st239
	tr437:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st239
	st239:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof239
		}
	st_case_239:
//line internal/php/scanner.go:8512
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr437
		case 2573:
			goto tr438
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr436
		}
		goto tr435
	tr430:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st240
	tr438:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st240
	st240:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof240
		}
	st_case_240:
//line internal/php/scanner.go:8592
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr429
		case 2573:
			goto tr430
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr428
		}
		goto tr439
	st241:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof241
		}
	st_case_241:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2399:
			goto tr173
		case 2427:
			goto tr174
		case 2570:
			goto tr429
		case 2573:
			goto tr430
		case 2655:
			goto tr440
		case 2683:
			goto tr441
		}
		switch {
		case _widec < 2625:
			switch {
			case _widec < 2401:
				if 2369 <= _widec && _widec <= 2394 {
					goto tr173
				}
			case _widec > 2426:
				switch {
				case _widec > 2559:
					if 2560 <= _widec && _widec <= 2624 {
						goto tr428
					}
				case _widec >= 2432:
					goto tr173
				}
			default:
				goto tr173
			}
		case _widec > 2650:
			switch {
			case _widec < 2657:
				if 2651 <= _widec && _widec <= 2656 {
					goto tr428
				}
			case _widec > 2682:
				switch {
				case _widec > 2687:
					if 2688 <= _widec && _widec <= 2815 {
						goto tr440
					}
				case _widec >= 2684:
					goto tr428
				}
			default:
				goto tr440
			}
		default:
			goto tr440
		}
		goto tr439
	st242:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof242
		}
	st_case_242:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2340:
			goto tr175
		case 2570:
			goto tr429
		case 2573:
			goto tr430
		case 2596:
			goto tr442
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr428
		}
		goto tr439
	tr444:
		lex.cs = 243
//line internal/php/scanner.rl:393
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_END_HEREDOC
			lex.cs = 139
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st243:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof243
		}
	st_case_243:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:8819
		if lex.data[(lex.p)] == 96 {
			goto st0
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto st244
	st244:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof244
		}
	st_case_244:
		if lex.data[(lex.p)] == 96 {
			goto tr444
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr444
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr444
				}
			case lex.data[(lex.p)] >= 91:
				goto tr444
			}
		default:
			goto tr444
		}
		goto st244
	tr176:
//line internal/php/scanner.rl:413
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st245
	tr177:
//line internal/php/scanner.rl:409
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 245
				goto _out
			}
		}
		goto st245
	tr179:
//line internal/php/scanner.rl:410
		lex.te = (lex.p) + 1
		{
			if lex.versionAtLeast(8, 0) {
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_NULLSAFE_OBJECT_OPERATOR
				{
					(lex.p)++
					lex.cs = 245
					goto _out
				}
			} else {
				lex.ungetCnt(4)
				{
					lex.top--
					lex.cs = lex.stack[lex.top]
					goto _again
				}
			}
		}
		goto st245
	tr445:
//line internal/php/scanner.rl:413
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st245
	tr450:
//line internal/php/scanner.rl:412
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('['))
			lex.call(245, 251)
			goto _out
		}
		goto st245
	tr451:
//line internal/php/scanner.rl:413
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st245
	tr453:
//line internal/php/scanner.rl:408
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 245
				goto _out
			}
		}
		goto st245
	tr456:
//line internal/php/scanner.rl:411
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 245
				goto _out
			}
		}
		goto st245
	st245:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof245
		}
	st_case_245:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:8916
		switch lex.data[(lex.p)] {
		case 36:
			goto st246
		case 45:
			goto tr447
		case 63:
			goto tr448
		case 91:
			goto tr450
		case 96:
			goto tr445
		}
		switch {
		case lex.data[(lex.p)] < 92:
			if lex.data[(lex.p)] <= 64 {
				goto tr445
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr445
			}
		default:
			goto tr445
		}
		goto st250
	st246:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof246
		}
	st_case_246:
		if lex.data[(lex.p)] == 96 {
			goto tr451
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr451
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr451
			}
		default:
			goto tr451
		}
		goto st247
	st247:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof247
		}
	st_case_247:
		if lex.data[(lex.p)] == 96 {
			goto tr453
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr453
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr453
				}
			case lex.data[(lex.p)] >= 91:
				goto tr453
			}
		default:
			goto tr453
		}
		goto st247
	tr447:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st248
	st248:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof248
		}
	st_case_248:
//line internal/php/scanner.go:8999
		if lex.data[(lex.p)] == 62 {
			goto st118
		}
		goto tr451
	st118:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof118
		}
	st_case_118:
		if lex.data[(lex.p)] == 96 {
			goto tr176
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr176
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr176
			}
		default:
			goto tr176
		}
		goto tr177
	tr448:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st249
	st249:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof249
		}
	st_case_249:
//line internal/php/scanner.go:9035
		if lex.data[(lex.p)] == 45 {
			goto st119
		}
		goto tr451
	st119:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof119
		}
	st_case_119:
		if lex.data[(lex.p)] == 62 {
			goto st120
		}
		goto tr176
	st120:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof120
		}
	st_case_120:
		if lex.data[(lex.p)] == 96 {
			goto tr176
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr176
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr176
			}
		default:
			goto tr176
		}
		goto tr179
	st250:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof250
		}
	st_case_250:
		if lex.data[(lex.p)] == 96 {
			goto tr456
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr456
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr456
				}
			case lex.data[(lex.p)] >= 91:
				goto tr456
			}
		default:
			goto tr456
		}
		goto st250
	tr180:
//line internal/php/scanner.rl:417
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 251
				goto _out
			}
		}
		goto st251
	tr457:
//line internal/php/scanner.rl:423
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st251
	tr458:
//line internal/php/scanner.rl:420
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st251
	tr461:
//line internal/php/scanner.rl:421
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 251
				goto _out
			}
		}
		goto st251
	tr465:
//line internal/php/scanner.rl:422
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(']'))
			lex.ret(2)
			goto _out
		}
		goto st251
	tr466:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:420
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st251
	tr467:
//line internal/php/scanner.rl:423
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st251
	tr468:
//line internal/php/scanner.rl:421
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 251
				goto _out
			}
		}
		goto st251
	tr470:
//line internal/php/scanner.rl:418
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 251
				goto _out
			}
		}
		goto st251
	tr471:
//line internal/php/scanner.rl:417
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 251
				goto _out
			}
		}
		goto st251
	tr476:
//line internal/php/scanner.rl:419
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 251
				goto _out
			}
		}
		goto st251
	st251:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof251
		}
	st_case_251:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:9176
		switch lex.data[(lex.p)] {
		case 10:
			goto tr459
		case 13:
			goto tr460
		case 32:
			goto tr458
		case 33:
			goto tr461
		case 35:
			goto tr458
		case 36:
			goto st254
		case 37:
			goto tr461
		case 39:
			goto tr458
		case 48:
			goto tr463
		case 92:
			goto tr458
		case 93:
			goto tr465
		case 96:
			goto tr457
		case 124:
			goto tr461
		case 126:
			goto tr461
		}
		switch {
		case lex.data[(lex.p)] < 40:
			switch {
			case lex.data[(lex.p)] < 9:
				if lex.data[(lex.p)] <= 8 {
					goto tr457
				}
			case lex.data[(lex.p)] > 12:
				if 14 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 38 {
					goto tr457
				}
			default:
				goto tr458
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 58:
				if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
					goto tr181
				}
			case lex.data[(lex.p)] > 64:
				switch {
				case lex.data[(lex.p)] > 94:
					if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
						goto tr457
					}
				case lex.data[(lex.p)] >= 91:
					goto tr461
				}
			default:
				goto tr461
			}
		default:
			goto tr461
		}
		goto st261
	tr459:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st252
	st252:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof252
		}
	st_case_252:
//line internal/php/scanner.go:9260
		goto tr466
	tr460:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st253
	st253:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof253
		}
	st_case_253:
//line internal/php/scanner.go:9279
		if lex.data[(lex.p)] == 10 {
			goto tr459
		}
		goto tr467
	st254:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof254
		}
	st_case_254:
		if lex.data[(lex.p)] == 96 {
			goto tr468
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr468
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr468
			}
		default:
			goto tr468
		}
		goto st255
	st255:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof255
		}
	st_case_255:
		if lex.data[(lex.p)] == 96 {
			goto tr470
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr470
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr470
				}
			case lex.data[(lex.p)] >= 91:
				goto tr470
			}
		default:
			goto tr470
		}
		goto st255
	tr463:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st256
	st256:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof256
		}
	st_case_256:
//line internal/php/scanner.go:9341
		switch lex.data[(lex.p)] {
		case 66:
			goto st122
		case 79:
			goto st123
		case 88:
			goto st124
		case 95:
			goto st121
		case 98:
			goto st122
		case 111:
			goto st123
		case 120:
			goto st124
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr181
		}
		goto tr471
	tr181:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st257
	st257:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof257
		}
	st_case_257:
//line internal/php/scanner.go:9372
		if lex.data[(lex.p)] == 95 {
			goto st121
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr181
		}
		goto tr471
	st121:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof121
		}
	st_case_121:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr181
		}
		goto tr180
	st122:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof122
		}
	st_case_122:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr182
		}
		goto tr180
	tr182:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st258
	st258:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof258
		}
	st_case_258:
//line internal/php/scanner.go:9408
		if lex.data[(lex.p)] == 95 {
			goto st122
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr182
		}
		goto tr471
	st123:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof123
		}
	st_case_123:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr183
		}
		goto tr180
	tr183:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st259
	st259:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof259
		}
	st_case_259:
//line internal/php/scanner.go:9435
		if lex.data[(lex.p)] == 95 {
			goto st123
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr183
		}
		goto tr471
	st124:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr184
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr184
			}
		default:
			goto tr184
		}
		goto tr180
	tr184:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st260
	st260:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof260
		}
	st_case_260:
//line internal/php/scanner.go:9471
		if lex.data[(lex.p)] == 95 {
			goto st124
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr184
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr184
			}
		default:
			goto tr184
		}
		goto tr471
	st261:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof261
		}
	st_case_261:
		if lex.data[(lex.p)] == 96 {
			goto tr476
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr476
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr476
				}
			case lex.data[(lex.p)] >= 91:
				goto tr476
			}
		default:
			goto tr476
		}
		goto st261
	tr185:
		lex.cs = 262
//line internal/php/scanner.rl:431
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	tr187:
		lex.cs = 262
//line internal/php/scanner.rl:430
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_STRING_VARNAME
			lex.cs = 139
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr477:
		lex.cs = 262
//line internal/php/scanner.rl:431
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	tr479:
		lex.cs = 262
//line internal/php/scanner.rl:431
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	st262:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof262
		}
	st_case_262:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:9550
		if lex.data[(lex.p)] == 96 {
			goto tr477
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr477
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr477
			}
		default:
			goto tr477
		}
		goto tr478
	tr478:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st263
	st263:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof263
		}
	st_case_263:
//line internal/php/scanner.go:9577
		switch lex.data[(lex.p)] {
		case 91:
			goto tr187
		case 96:
			goto tr479
		case 125:
			goto tr187
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr479
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr479
				}
			case lex.data[(lex.p)] >= 92:
				goto tr479
			}
		default:
			goto tr479
		}
		goto st125
	st125:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof125
		}
	st_case_125:
		switch lex.data[(lex.p)] {
		case 91:
			goto tr187
		case 96:
			goto tr185
		case 125:
			goto tr187
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr185
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr185
				}
			case lex.data[(lex.p)] >= 92:
				goto tr185
			}
		default:
			goto tr185
		}
		goto st125
	tr188:
//line internal/php/scanner.rl:435
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st264
	tr480:
		lex.cs = 264
//line internal/php/scanner.rl:437
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	tr483:
		lex.cs = 264
//line internal/php/scanner.rl:436
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('('))
			lex.cs = 268
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr484:
//line internal/php/scanner.rl:435
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st264
	tr486:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:435
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st264
	tr490:
		lex.cs = 264
//line internal/php/scanner.rl:437
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	st264:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof264
		}
	st_case_264:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:9684
		switch lex.data[(lex.p)] {
		case 10:
			goto tr189
		case 13:
			goto tr482
		case 32:
			goto tr481
		case 40:
			goto tr483
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr481
		}
		goto tr480
	tr487:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

		goto st265
	tr481:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st265
	st265:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof265
		}
	st_case_265:
//line internal/php/scanner.go:9716
		switch lex.data[(lex.p)] {
		case 10:
			goto tr189
		case 13:
			goto tr485
		case 32:
			goto tr481
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr481
		}
		goto tr484
	tr189:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st266
	tr488:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st266
	st266:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof266
		}
	st_case_266:
//line internal/php/scanner.go:9766
		switch lex.data[(lex.p)] {
		case 10:
			goto tr488
		case 13:
			goto tr489
		case 32:
			goto tr487
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr487
		}
		goto tr486
	tr485:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st126
	tr489:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st126
	st126:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof126
		}
	st_case_126:
//line internal/php/scanner.go:9810
		if lex.data[(lex.p)] == 10 {
			goto tr189
		}
		goto tr188
	tr482:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st267
	st267:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof267
		}
	st_case_267:
//line internal/php/scanner.go:9832
		if lex.data[(lex.p)] == 10 {
			goto tr189
		}
		goto tr490
	tr190:
//line internal/php/scanner.rl:441
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st268
	tr491:
		lex.cs = 268
//line internal/php/scanner.rl:443
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	tr494:
		lex.cs = 268
//line internal/php/scanner.rl:442
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(')'))
			lex.cs = 272
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr495:
//line internal/php/scanner.rl:441
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st268
	tr497:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:441
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st268
	tr501:
		lex.cs = 268
//line internal/php/scanner.rl:443
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	st268:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof268
		}
	st_case_268:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:9886
		switch lex.data[(lex.p)] {
		case 10:
			goto tr191
		case 13:
			goto tr493
		case 32:
			goto tr492
		case 41:
			goto tr494
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr492
		}
		goto tr491
	tr498:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

		goto st269
	tr492:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st269
	st269:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof269
		}
	st_case_269:
//line internal/php/scanner.go:9918
		switch lex.data[(lex.p)] {
		case 10:
			goto tr191
		case 13:
			goto tr496
		case 32:
			goto tr492
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr492
		}
		goto tr495
	tr191:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st270
	tr499:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st270
	st270:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof270
		}
	st_case_270:
//line internal/php/scanner.go:9968
		switch lex.data[(lex.p)] {
		case 10:
			goto tr499
		case 13:
			goto tr500
		case 32:
			goto tr498
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr498
		}
		goto tr497
	tr496:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st127
	tr500:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st127
	st127:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof127
		}
	st_case_127:
//line internal/php/scanner.go:10012
		if lex.data[(lex.p)] == 10 {
			goto tr191
		}
		goto tr190
	tr493:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st271
	st271:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof271
		}
	st_case_271:
//line internal/php/scanner.go:10034
		if lex.data[(lex.p)] == 10 {
			goto tr191
		}
		goto tr501
	tr192:
//line internal/php/scanner.rl:447
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st272
	tr502:
		lex.cs = 272
//line internal/php/scanner.rl:449
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	tr505:
		lex.cs = 272
//line internal/php/scanner.rl:448
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 276
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr506:
//line internal/php/scanner.rl:447
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st272
	tr508:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:447
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st272
	tr512:
		lex.cs = 272
//line internal/php/scanner.rl:449
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 139
		}
		goto _again
	st272:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof272
		}
	st_case_272:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:10088
		switch lex.data[(lex.p)] {
		case 10:
			goto tr193
		case 13:
			goto tr504
		case 32:
			goto tr503
		case 59:
			goto tr505
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr503
		}
		goto tr502
	tr509:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

		goto st273
	tr503:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st273
	st273:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof273
		}
	st_case_273:
//line internal/php/scanner.go:10120
		switch lex.data[(lex.p)] {
		case 10:
			goto tr193
		case 13:
			goto tr507
		case 32:
			goto tr503
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr503
		}
		goto tr506
	tr193:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st274
	tr510:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st274
	st274:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof274
		}
	st_case_274:
//line internal/php/scanner.go:10170
		switch lex.data[(lex.p)] {
		case 10:
			goto tr510
		case 13:
			goto tr511
		case 32:
			goto tr509
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr509
		}
		goto tr508
	tr507:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st128
	tr511:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st128
	st128:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof128
		}
	st_case_128:
//line internal/php/scanner.go:10214
		if lex.data[(lex.p)] == 10 {
			goto tr193
		}
		goto tr192
	tr504:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st275
	st275:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof275
		}
	st_case_275:
//line internal/php/scanner.go:10236
		if lex.data[(lex.p)] == 10 {
			goto tr193
		}
		goto tr512
	tr516:
//line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 128:
			{
				(lex.p) = (lex.te) - 1
				lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
			}
		}

		goto st276
	tr517:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:453
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
		}
		goto st276
	tr521:
//line internal/php/scanner.rl:453
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
		}
		goto st276
	st276:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof276
		}
	st_case_276:
//line NONE:1
		lex.ts = (lex.p)

//line internal/php/scanner.go:10280
		switch lex.data[(lex.p)] {
		case 10:
			goto tr514
		case 13:
			goto tr515
		}
		goto tr513
	tr513:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:453
		lex.act = 128
		goto st277
	tr518:
//line NONE:1
		lex.te = (lex.p) + 1

//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:453
		lex.act = 128
		goto st277
	st277:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof277
		}
	st_case_277:
//line internal/php/scanner.go:10309
		switch lex.data[(lex.p)] {
		case 10:
			goto tr514
		case 13:
			goto tr515
		}
		goto tr513
	tr514:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st278
	tr519:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st278
	st278:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof278
		}
	st_case_278:
//line internal/php/scanner.go:10348
		switch lex.data[(lex.p)] {
		case 10:
			goto tr519
		case 13:
			goto tr520
		}
		goto tr518
	tr515:
//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st279
	tr520:
//line internal/php/scanner.rl:53

//line internal/php/scanner.rl:37

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st279
	st279:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof279
		}
	st_case_279:
//line internal/php/scanner.go:10387
		switch lex.data[(lex.p)] {
		case 10:
			goto tr514
		case 13:
			goto tr515
		}
		goto tr513
	st_out:
	_test_eof129:
		lex.cs = 129
		goto _test_eof
	_test_eof130:
		lex.cs = 130
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof131:
		lex.cs = 131
		goto _test_eof
	_test_eof132:
		lex.cs = 132
		goto _test_eof
	_test_eof133:
		lex.cs = 133
		goto _test_eof
	_test_eof134:
		lex.cs = 134
		goto _test_eof
	_test_eof135:
		lex.cs = 135
		goto _test_eof
	_test_eof136:
		lex.cs = 136
		goto _test_eof
	_test_eof137:
		lex.cs = 137
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof138:
		lex.cs = 138
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof139:
		lex.cs = 139
		goto _test_eof
	_test_eof140:
		lex.cs = 140
		goto _test_eof
	_test_eof141:
		lex.cs = 141
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof142:
		lex.cs = 142
		goto _test_eof
	_test_eof143:
		lex.cs = 143
		goto _test_eof
	_test_eof144:
		lex.cs = 144
		goto _test_eof
	_test_eof145:
		lex.cs = 145
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof146:
		lex.cs = 146
		goto _test_eof
	_test_eof147:
		lex.cs = 147
		goto _test_eof
	_test_eof148:
		lex.cs = 148
		goto _test_eof
	_test_eof149:
		lex.cs = 149
		goto _test_eof
	_test_eof150:
		lex.cs = 150
		goto _test_eof
	_test_eof151:
		lex.cs = 151
		goto _test_eof
	_test_eof152:
		lex.cs = 152
		goto _test_eof
	_test_eof153:
		lex.cs = 153
		goto _test_eof
	_test_eof154:
		lex.cs = 154
		goto _test_eof
	_test_eof155:
		lex.cs = 155
		goto _test_eof
	_test_eof156:
		lex.cs = 156
		goto _test_eof
	_test_eof157:
		lex.cs = 157
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof
	_test_eof158:
		lex.cs = 158
		goto _test_eof
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof159:
		lex.cs = 159
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof160:
		lex.cs = 160
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof18:
		lex.cs = 18
		goto _test_eof
	_test_eof19:
		lex.cs = 19
		goto _test_eof
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof22:
		lex.cs = 22
		goto _test_eof
	_test_eof23:
		lex.cs = 23
		goto _test_eof
	_test_eof24:
		lex.cs = 24
		goto _test_eof
	_test_eof25:
		lex.cs = 25
		goto _test_eof
	_test_eof26:
		lex.cs = 26
		goto _test_eof
	_test_eof27:
		lex.cs = 27
		goto _test_eof
	_test_eof28:
		lex.cs = 28
		goto _test_eof
	_test_eof29:
		lex.cs = 29
		goto _test_eof
	_test_eof30:
		lex.cs = 30
		goto _test_eof
	_test_eof31:
		lex.cs = 31
		goto _test_eof
	_test_eof32:
		lex.cs = 32
		goto _test_eof
	_test_eof33:
		lex.cs = 33
		goto _test_eof
	_test_eof34:
		lex.cs = 34
		goto _test_eof
	_test_eof35:
		lex.cs = 35
		goto _test_eof
	_test_eof36:
		lex.cs = 36
		goto _test_eof
	_test_eof37:
		lex.cs = 37
		goto _test_eof
	_test_eof38:
		lex.cs = 38
		goto _test_eof
	_test_eof39:
		lex.cs = 39
		goto _test_eof
	_test_eof40:
		lex.cs = 40
		goto _test_eof
	_test_eof41:
		lex.cs = 41
		goto _test_eof
	_test_eof42:
		lex.cs = 42
		goto _test_eof
	_test_eof43:
		lex.cs = 43
		goto _test_eof
	_test_eof44:
		lex.cs = 44
		goto _test_eof
	_test_eof45:
		lex.cs = 45
		goto _test_eof
	_test_eof46:
		lex.cs = 46
		goto _test_eof
	_test_eof47:
		lex.cs = 47
		goto _test_eof
	_test_eof48:
		lex.cs = 48
		goto _test_eof
	_test_eof49:
		lex.cs = 49
		goto _test_eof
	_test_eof50:
		lex.cs = 50
		goto _test_eof
	_test_eof51:
		lex.cs = 51
		goto _test_eof
	_test_eof52:
		lex.cs = 52
		goto _test_eof
	_test_eof53:
		lex.cs = 53
		goto _test_eof
	_test_eof54:
		lex.cs = 54
		goto _test_eof
	_test_eof55:
		lex.cs = 55
		goto _test_eof
	_test_eof56:
		lex.cs = 56
		goto _test_eof
	_test_eof57:
		lex.cs = 57
		goto _test_eof
	_test_eof58:
		lex.cs = 58
		goto _test_eof
	_test_eof59:
		lex.cs = 59
		goto _test_eof
	_test_eof60:
		lex.cs = 60
		goto _test_eof
	_test_eof61:
		lex.cs = 61
		goto _test_eof
	_test_eof62:
		lex.cs = 62
		goto _test_eof
	_test_eof63:
		lex.cs = 63
		goto _test_eof
	_test_eof64:
		lex.cs = 64
		goto _test_eof
	_test_eof65:
		lex.cs = 65
		goto _test_eof
	_test_eof66:
		lex.cs = 66
		goto _test_eof
	_test_eof67:
		lex.cs = 67
		goto _test_eof
	_test_eof68:
		lex.cs = 68
		goto _test_eof
	_test_eof69:
		lex.cs = 69
		goto _test_eof
	_test_eof70:
		lex.cs = 70
		goto _test_eof
	_test_eof71:
		lex.cs = 71
		goto _test_eof
	_test_eof72:
		lex.cs = 72
		goto _test_eof
	_test_eof73:
		lex.cs = 73
		goto _test_eof
	_test_eof161:
		lex.cs = 161
		goto _test_eof
	_test_eof162:
		lex.cs = 162
		goto _test_eof
	_test_eof163:
		lex.cs = 163
		goto _test_eof
	_test_eof164:
		lex.cs = 164
		goto _test_eof
	_test_eof165:
		lex.cs = 165
		goto _test_eof
	_test_eof74:
		lex.cs = 74
		goto _test_eof
	_test_eof166:
		lex.cs = 166
		goto _test_eof
	_test_eof75:
		lex.cs = 75
		goto _test_eof
	_test_eof76:
		lex.cs = 76
		goto _test_eof
	_test_eof167:
		lex.cs = 167
		goto _test_eof
	_test_eof77:
		lex.cs = 77
		goto _test_eof
	_test_eof168:
		lex.cs = 168
		goto _test_eof
	_test_eof78:
		lex.cs = 78
		goto _test_eof
	_test_eof79:
		lex.cs = 79
		goto _test_eof
	_test_eof80:
		lex.cs = 80
		goto _test_eof
	_test_eof169:
		lex.cs = 169
		goto _test_eof
	_test_eof170:
		lex.cs = 170
		goto _test_eof
	_test_eof171:
		lex.cs = 171
		goto _test_eof
	_test_eof81:
		lex.cs = 81
		goto _test_eof
	_test_eof82:
		lex.cs = 82
		goto _test_eof
	_test_eof172:
		lex.cs = 172
		goto _test_eof
	_test_eof83:
		lex.cs = 83
		goto _test_eof
	_test_eof173:
		lex.cs = 173
		goto _test_eof
	_test_eof84:
		lex.cs = 84
		goto _test_eof
	_test_eof174:
		lex.cs = 174
		goto _test_eof
	_test_eof175:
		lex.cs = 175
		goto _test_eof
	_test_eof176:
		lex.cs = 176
		goto _test_eof
	_test_eof85:
		lex.cs = 85
		goto _test_eof
	_test_eof86:
		lex.cs = 86
		goto _test_eof
	_test_eof87:
		lex.cs = 87
		goto _test_eof
	_test_eof88:
		lex.cs = 88
		goto _test_eof
	_test_eof177:
		lex.cs = 177
		goto _test_eof
	_test_eof178:
		lex.cs = 178
		goto _test_eof
	_test_eof89:
		lex.cs = 89
		goto _test_eof
	_test_eof179:
		lex.cs = 179
		goto _test_eof
	_test_eof180:
		lex.cs = 180
		goto _test_eof
	_test_eof90:
		lex.cs = 90
		goto _test_eof
	_test_eof91:
		lex.cs = 91
		goto _test_eof
	_test_eof92:
		lex.cs = 92
		goto _test_eof
	_test_eof93:
		lex.cs = 93
		goto _test_eof
	_test_eof181:
		lex.cs = 181
		goto _test_eof
	_test_eof94:
		lex.cs = 94
		goto _test_eof
	_test_eof95:
		lex.cs = 95
		goto _test_eof
	_test_eof96:
		lex.cs = 96
		goto _test_eof
	_test_eof97:
		lex.cs = 97
		goto _test_eof
	_test_eof182:
		lex.cs = 182
		goto _test_eof
	_test_eof183:
		lex.cs = 183
		goto _test_eof
	_test_eof184:
		lex.cs = 184
		goto _test_eof
	_test_eof185:
		lex.cs = 185
		goto _test_eof
	_test_eof186:
		lex.cs = 186
		goto _test_eof
	_test_eof187:
		lex.cs = 187
		goto _test_eof
	_test_eof98:
		lex.cs = 98
		goto _test_eof
	_test_eof188:
		lex.cs = 188
		goto _test_eof
	_test_eof189:
		lex.cs = 189
		goto _test_eof
	_test_eof99:
		lex.cs = 99
		goto _test_eof
	_test_eof190:
		lex.cs = 190
		goto _test_eof
	_test_eof191:
		lex.cs = 191
		goto _test_eof
	_test_eof100:
		lex.cs = 100
		goto _test_eof
	_test_eof192:
		lex.cs = 192
		goto _test_eof
	_test_eof193:
		lex.cs = 193
		goto _test_eof
	_test_eof101:
		lex.cs = 101
		goto _test_eof
	_test_eof102:
		lex.cs = 102
		goto _test_eof
	_test_eof194:
		lex.cs = 194
		goto _test_eof
	_test_eof195:
		lex.cs = 195
		goto _test_eof
	_test_eof196:
		lex.cs = 196
		goto _test_eof
	_test_eof197:
		lex.cs = 197
		goto _test_eof
	_test_eof198:
		lex.cs = 198
		goto _test_eof
	_test_eof199:
		lex.cs = 199
		goto _test_eof
	_test_eof200:
		lex.cs = 200
		goto _test_eof
	_test_eof201:
		lex.cs = 201
		goto _test_eof
	_test_eof202:
		lex.cs = 202
		goto _test_eof
	_test_eof103:
		lex.cs = 103
		goto _test_eof
	_test_eof203:
		lex.cs = 203
		goto _test_eof
	_test_eof204:
		lex.cs = 204
		goto _test_eof
	_test_eof205:
		lex.cs = 205
		goto _test_eof
	_test_eof206:
		lex.cs = 206
		goto _test_eof
	_test_eof207:
		lex.cs = 207
		goto _test_eof
	_test_eof208:
		lex.cs = 208
		goto _test_eof
	_test_eof104:
		lex.cs = 104
		goto _test_eof
	_test_eof105:
		lex.cs = 105
		goto _test_eof
	_test_eof106:
		lex.cs = 106
		goto _test_eof
	_test_eof107:
		lex.cs = 107
		goto _test_eof
	_test_eof108:
		lex.cs = 108
		goto _test_eof
	_test_eof109:
		lex.cs = 109
		goto _test_eof
	_test_eof209:
		lex.cs = 209
		goto _test_eof
	_test_eof210:
		lex.cs = 210
		goto _test_eof
	_test_eof110:
		lex.cs = 110
		goto _test_eof
	_test_eof211:
		lex.cs = 211
		goto _test_eof
	_test_eof212:
		lex.cs = 212
		goto _test_eof
	_test_eof213:
		lex.cs = 213
		goto _test_eof
	_test_eof214:
		lex.cs = 214
		goto _test_eof
	_test_eof215:
		lex.cs = 215
		goto _test_eof
	_test_eof111:
		lex.cs = 111
		goto _test_eof
	_test_eof216:
		lex.cs = 216
		goto _test_eof
	_test_eof217:
		lex.cs = 217
		goto _test_eof
	_test_eof218:
		lex.cs = 218
		goto _test_eof
	_test_eof112:
		lex.cs = 112
		goto _test_eof
	_test_eof219:
		lex.cs = 219
		goto _test_eof
	_test_eof220:
		lex.cs = 220
		goto _test_eof
	_test_eof221:
		lex.cs = 221
		goto _test_eof
	_test_eof222:
		lex.cs = 222
		goto _test_eof
	_test_eof223:
		lex.cs = 223
		goto _test_eof
	_test_eof224:
		lex.cs = 224
		goto _test_eof
	_test_eof225:
		lex.cs = 225
		goto _test_eof
	_test_eof113:
		lex.cs = 113
		goto _test_eof
	_test_eof226:
		lex.cs = 226
		goto _test_eof
	_test_eof227:
		lex.cs = 227
		goto _test_eof
	_test_eof228:
		lex.cs = 228
		goto _test_eof
	_test_eof229:
		lex.cs = 229
		goto _test_eof
	_test_eof230:
		lex.cs = 230
		goto _test_eof
	_test_eof231:
		lex.cs = 231
		goto _test_eof
	_test_eof114:
		lex.cs = 114
		goto _test_eof
	_test_eof115:
		lex.cs = 115
		goto _test_eof
	_test_eof232:
		lex.cs = 232
		goto _test_eof
	_test_eof233:
		lex.cs = 233
		goto _test_eof
	_test_eof234:
		lex.cs = 234
		goto _test_eof
	_test_eof235:
		lex.cs = 235
		goto _test_eof
	_test_eof236:
		lex.cs = 236
		goto _test_eof
	_test_eof237:
		lex.cs = 237
		goto _test_eof
	_test_eof116:
		lex.cs = 116
		goto _test_eof
	_test_eof117:
		lex.cs = 117
		goto _test_eof
	_test_eof238:
		lex.cs = 238
		goto _test_eof
	_test_eof239:
		lex.cs = 239
		goto _test_eof
	_test_eof240:
		lex.cs = 240
		goto _test_eof
	_test_eof241:
		lex.cs = 241
		goto _test_eof
	_test_eof242:
		lex.cs = 242
		goto _test_eof
	_test_eof243:
		lex.cs = 243
		goto _test_eof
	_test_eof244:
		lex.cs = 244
		goto _test_eof
	_test_eof245:
		lex.cs = 245
		goto _test_eof
	_test_eof246:
		lex.cs = 246
		goto _test_eof
	_test_eof247:
		lex.cs = 247
		goto _test_eof
	_test_eof248:
		lex.cs = 248
		goto _test_eof
	_test_eof118:
		lex.cs = 118
		goto _test_eof
	_test_eof249:
		lex.cs = 249
		goto _test_eof
	_test_eof119:
		lex.cs = 119
		goto _test_eof
	_test_eof120:
		lex.cs = 120
		goto _test_eof
	_test_eof250:
		lex.cs = 250
		goto _test_eof
	_test_eof251:
		lex.cs = 251
		goto _test_eof
	_test_eof252:
		lex.cs = 252
		goto _test_eof
	_test_eof253:
		lex.cs = 253
		goto _test_eof
	_test_eof254:
		lex.cs = 254
		goto _test_eof
	_test_eof255:
		lex.cs = 255
		goto _test_eof
	_test_eof256:
		lex.cs = 256
		goto _test_eof
	_test_eof257:
		lex.cs = 257
		goto _test_eof
	_test_eof121:
		lex.cs = 121
		goto _test_eof
	_test_eof122:
		lex.cs = 122
		goto _test_eof
	_test_eof258:
		lex.cs = 258
		goto _test_eof
	_test_eof123:
		lex.cs = 123
		goto _test_eof
	_test_eof259:
		lex.cs = 259
		goto _test_eof
	_test_eof124:
		lex.cs = 124
		goto _test_eof
	_test_eof260:
		lex.cs = 260
		goto _test_eof
	_test_eof261:
		lex.cs = 261
		goto _test_eof
	_test_eof262:
		lex.cs = 262
		goto _test_eof
	_test_eof263:
		lex.cs = 263
		goto _test_eof
	_test_eof125:
		lex.cs = 125
		goto _test_eof
	_test_eof264:
		lex.cs = 264
		goto _test_eof
	_test_eof265:
		lex.cs = 265
		goto _test_eof
	_test_eof266:
		lex.cs = 266
		goto _test_eof
	_test_eof126:
		lex.cs = 126
		goto _test_eof
	_test_eof267:
		lex.cs = 267
		goto _test_eof
	_test_eof268:
		lex.cs = 268
		goto _test_eof
	_test_eof269:
		lex.cs = 269
		goto _test_eof
	_test_eof270:
		lex.cs = 270
		goto _test_eof
	_test_eof127:
		lex.cs = 127
		goto _test_eof
	_test_eof271:
		lex.cs = 271
		goto _test_eof
	_test_eof272:
		lex.cs = 272
		goto _test_eof
	_test_eof273:
		lex.cs = 273
		goto _test_eof
	_test_eof274:
		lex.cs = 274
		goto _test_eof
	_test_eof128:
		lex.cs = 128
		goto _test_eof
	_test_eof275:
		lex.cs = 275
		goto _test_eof
	_test_eof276:
		lex.cs = 276
		goto _test_eof
	_test_eof277:
		lex.cs = 277
		goto _test_eof
	_test_eof278:
		lex.cs = 278
		goto _test_eof
	_test_eof279:
		lex.cs = 279
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 130:
				goto tr196
			case 1:
				goto tr0
			case 131:
				goto tr197
			case 133:
				goto tr202
			case 134:
				goto tr204
			case 135:
				goto tr202
			case 136:
				goto tr202
			case 137:
				goto tr210
			case 2:
				goto tr4
			case 3:
				goto tr4
			case 4:
				goto tr4
			case 138:
				goto tr213
			case 5:
				goto tr4
			case 140:
				goto tr248
			case 141:
				goto tr250
			case 6:
				goto tr10
			case 142:
				goto tr254
			case 143:
				goto tr255
			case 144:
				goto tr257
			case 145:
				goto tr259
			case 7:
				goto tr12
			case 8:
				goto tr12
			case 9:
				goto tr12
			case 10:
				goto tr12
			case 146:
				goto tr260
			case 147:
				goto tr265
			case 148:
				goto tr268
			case 149:
				goto tr272
			case 150:
				goto tr265
			case 151:
				goto tr255
			case 152:
				goto tr275
			case 153:
				goto tr255
			case 154:
				goto tr254
			case 155:
				goto tr285
			case 156:
				goto tr285
			case 157:
				goto tr285
			case 11:
				goto tr19
			case 158:
				goto tr285
			case 12:
				goto tr19
			case 159:
				goto tr254
			case 13:
				goto tr27
			case 14:
				goto tr27
			case 160:
				goto tr255
			case 15:
				goto tr31
			case 16:
				goto tr31
			case 17:
				goto tr31
			case 18:
				goto tr31
			case 19:
				goto tr31
			case 20:
				goto tr31
			case 21:
				goto tr31
			case 22:
				goto tr31
			case 23:
				goto tr31
			case 24:
				goto tr31
			case 25:
				goto tr31
			case 26:
				goto tr31
			case 27:
				goto tr31
			case 28:
				goto tr31
			case 29:
				goto tr31
			case 30:
				goto tr31
			case 31:
				goto tr31
			case 32:
				goto tr31
			case 33:
				goto tr31
			case 34:
				goto tr31
			case 35:
				goto tr31
			case 36:
				goto tr31
			case 37:
				goto tr31
			case 38:
				goto tr31
			case 39:
				goto tr31
			case 40:
				goto tr31
			case 41:
				goto tr31
			case 42:
				goto tr31
			case 43:
				goto tr31
			case 44:
				goto tr31
			case 45:
				goto tr31
			case 46:
				goto tr31
			case 47:
				goto tr31
			case 48:
				goto tr31
			case 49:
				goto tr31
			case 50:
				goto tr31
			case 51:
				goto tr31
			case 52:
				goto tr31
			case 53:
				goto tr31
			case 54:
				goto tr31
			case 55:
				goto tr31
			case 56:
				goto tr31
			case 57:
				goto tr31
			case 58:
				goto tr31
			case 59:
				goto tr31
			case 60:
				goto tr31
			case 61:
				goto tr31
			case 62:
				goto tr31
			case 63:
				goto tr31
			case 64:
				goto tr31
			case 65:
				goto tr31
			case 66:
				goto tr31
			case 67:
				goto tr31
			case 68:
				goto tr31
			case 69:
				goto tr31
			case 70:
				goto tr31
			case 71:
				goto tr31
			case 72:
				goto tr31
			case 73:
				goto tr31
			case 161:
				goto tr255
			case 162:
				goto tr290
			case 163:
				goto tr255
			case 164:
				goto tr255
			case 165:
				goto tr255
			case 74:
				goto tr31
			case 166:
				goto tr299
			case 75:
				goto tr12
			case 76:
				goto tr12
			case 167:
				goto tr299
			case 77:
				goto tr103
			case 168:
				goto tr255
			case 78:
				goto tr31
			case 79:
				goto tr31
			case 80:
				goto tr31
			case 169:
				goto tr303
			case 170:
				goto tr299
			case 171:
				goto tr303
			case 81:
				goto tr114
			case 82:
				goto tr12
			case 172:
				goto tr309
			case 83:
				goto tr12
			case 173:
				goto tr310
			case 84:
				goto tr12
			case 174:
				goto tr311
			case 175:
				goto tr255
			case 176:
				goto tr255
			case 85:
				goto tr31
			case 86:
				goto tr31
			case 87:
				goto tr31
			case 88:
				goto tr31
			case 177:
				goto tr313
			case 178:
				goto tr315
			case 89:
				goto tr128
			case 179:
				goto tr255
			case 180:
				goto tr319
			case 90:
				goto tr12
			case 91:
				goto tr12
			case 92:
				goto tr12
			case 93:
				goto tr12
			case 181:
				goto tr321
			case 94:
				goto tr12
			case 95:
				goto tr12
			case 96:
				goto tr12
			case 97:
				goto tr12
			case 182:
				goto tr322
			case 183:
				goto tr255
			case 184:
				goto tr326
			case 185:
				goto tr255
			case 186:
				goto tr330
			case 187:
				goto tr255
			case 98:
				goto tr31
			case 188:
				goto tr335
			case 189:
				goto tr337
			case 99:
				goto tr145
			case 190:
				goto tr338
			case 191:
				goto tr340
			case 100:
				goto tr12
			case 192:
				goto tr342
			case 193:
				goto tr340
			case 101:
				goto tr148
			case 102:
				goto tr148
			case 194:
				goto tr340
			case 195:
				goto tr340
			case 196:
				goto tr340
			case 197:
				goto tr340
			case 198:
				goto tr340
			case 199:
				goto tr340
			case 200:
				goto tr340
			case 201:
				goto tr340
			case 202:
				goto tr340
			case 103:
				goto tr12
			case 203:
				goto tr353
			case 204:
				goto tr340
			case 205:
				goto tr340
			case 206:
				goto tr340
			case 207:
				goto tr340
			case 208:
				goto tr340
			case 104:
				goto tr148
			case 105:
				goto tr148
			case 106:
				goto tr148
			case 107:
				goto tr148
			case 108:
				goto tr148
			case 109:
				goto tr148
			case 209:
				goto tr358
			case 210:
				goto tr359
			case 110:
				goto tr162
			case 211:
				goto tr255
			case 212:
				goto tr255
			case 214:
				goto tr371
			case 215:
				goto tr373
			case 111:
				goto tr164
			case 216:
				goto tr377
			case 217:
				goto tr377
			case 218:
				goto tr377
			case 112:
				goto tr166
			case 219:
				goto tr380
			case 221:
				goto tr384
			case 222:
				goto tr385
			case 223:
				goto tr389
			case 225:
				goto tr397
			case 226:
				goto tr399
			case 227:
				goto tr400
			case 228:
				goto tr404
			case 229:
				goto tr397
			case 230:
				goto tr404
			case 232:
				goto tr416
			case 233:
				goto tr417
			case 234:
				goto tr421
			case 235:
				goto tr421
			case 236:
				goto tr421
			case 238:
				goto tr434
			case 239:
				goto tr435
			case 240:
				goto tr439
			case 241:
				goto tr439
			case 242:
				goto tr439
			case 244:
				goto tr444
			case 246:
				goto tr451
			case 247:
				goto tr453
			case 248:
				goto tr451
			case 118:
				goto tr176
			case 249:
				goto tr451
			case 119:
				goto tr176
			case 120:
				goto tr176
			case 250:
				goto tr456
			case 252:
				goto tr466
			case 253:
				goto tr467
			case 254:
				goto tr468
			case 255:
				goto tr470
			case 256:
				goto tr471
			case 257:
				goto tr471
			case 121:
				goto tr180
			case 122:
				goto tr180
			case 258:
				goto tr471
			case 123:
				goto tr180
			case 259:
				goto tr471
			case 124:
				goto tr180
			case 260:
				goto tr471
			case 261:
				goto tr476
			case 263:
				goto tr479
			case 125:
				goto tr185
			case 265:
				goto tr484
			case 266:
				goto tr486
			case 126:
				goto tr188
			case 267:
				goto tr490
			case 269:
				goto tr495
			case 270:
				goto tr497
			case 127:
				goto tr190
			case 271:
				goto tr501
			case 273:
				goto tr506
			case 274:
				goto tr508
			case 128:
				goto tr192
			case 275:
				goto tr512
			case 277:
				goto tr516
			case 278:
				goto tr517
			case 279:
				goto tr521
			}
		}

	_out:
		{
		}
	}

//line internal/php/scanner.rl:457

	tkn.Value = lex.data[lex.ts:lex.te]
	tkn.ID = token.ID(tok)

	return tkn
}
