package parser

import (
  "fmt"
  "os"
  "sstat/inst"
  "sstat/lexer"
  "strings"
  "strconv"
)

var PcFound bool = false

type EbbaParser struct {
    tlexer    *lexer.Lexer
    file_name string
}

func (eparser *EbbaParser) CheckEOF() bool {
    var xtok = eparser.tlexer.Peek(1)
    if xtok.TokType == lexer.TOK_EOF {
      fmt.Printf("EOF st line %d\n", xtok.LineNum)
      return true
    }
    return false
}

func (eparser *EbbaParser) getCore() uint64 {
    var tlexer = eparser.tlexer
    tlexer.Match(lexer.TOK_LT, "Expected a '<' symbol\n")
    var xtok = tlexer.GetToken()

    if !string.HasPrefix(string(xtok.Value), "core"){
      fmt.Println("Expected 'core' in line")
      fmt.Println("Found: ")
      lexer.PrintTok(xtok)
      os.Exit(1)
    }
    
    core_str := string(xtok.Value)[4:]
    core, err := strconv.ParseUint(core_str, 10, 64)
    if err != nil {
      fmt.Println("Expected a decimal number after 'core', found ", core_str)
      os.Exit(1)
    }
    tlexer.Match(lexer.TOK_GT, "Expected a '>' symbol\n")
    return core

}

func (eparser *EbbaParser) handle_reset_instr(core uint64, instr *inst.Inst){
    var tlexer = eparser.tlexer
    var xtok *lexer.Token = tlexer.Peek(1)
    if xtok.TokType == lexer.TOK_XREG {
      tlexer.Match(lexer.TOK_XREG, "Expected an X register\n")
    }
    else if xtok.TokType == lexer.TOK_CSR {
      tlexer.Match(lexer.TOK_CSR, "Expected a CSR register\n")
    }
    else{
      tlexer.Match(lexer.TOK_ID, "Expected an ID\n")
    }

    tlexer.Match(lexer.TOK_LT, "Expected an '<' Symbol\n")
    tlexer.Match(lexer.TOK_MINUS_SIGN, "Expected an '-' Sign\n")
    tlexer.Match(lexer.HEX_NUMBER, "Expected an Hexadecimal Number\n")
    
    xtok = tlexer.GetToken()
    if xtok.TokType != lexer.TOK_NEWLINE && xtok.TokType != lexer.TOK_SEMI_COLON {
      fmt.Println("Expected New Line or a Semicolon")
      fmt.Println("Found: ")
      lexer.PrintTok(xtok)
      os.Exit(1)
    }
    
    tlexer.ConsumeTillNewLine()

}

func (eparser *EbbaParser) handle_instruction_lines (core uint64, instr *inst.Inst){
	var tlexer = eparser.tlexer
	var xtok = *lexer.Token
	var err error

	tlexer.Match(lexer.TOK_L_PAREN, "Expected token '('\n")

	xtok = tlexer.GetToken()

	instrNum, err := strconv.ParseUint(string(xtok.Value), 10, 64)

	if err != nil {
		fmt.Println("Expected a decimal number")
		fmt.Println("Found: ")
		lexer.PrintTok(xtok)
		os.Exit(1)
	}

	tlexer.Match(lexer.TOK_R_PAREN, "Expected token ')'\n")
	tlexer.Match(lexer.TOK_MINUS_SIGN, "Expected token '-'\n")
	tlexer.Match(lexer.TOK_GT, "Expected token '>'\n")

	xtok = tlexer.GetToken()

	instr.OpCode, err = strconv.ParseUint(string(xtok.Value[2:]), 16, 64)
	if err != nil {
		fmt.Printf("Unable to convert value to uint64 : %s \n", xtok.Value)
		os.Exit(1)
	}
	tlexer.Match(lexer.TOK_SEMI_COLON, "Expected a semicolon\n")
	eparser.ebba_instr(instr)

}

func (eparser *EbbaParser) GetInst() (uint64, *inst.Inst){
    var instr = new(inst.Inst)
    var tlexer = eparser.tlexer
    var ptok *lexer.Token
    var xtok *lexer.Token
    var err error
    var core uint64
    var instructionParsed bool = false
    var instLineSeen bool = false

    for true {
        if instructionParsed {
            break
        }
        ptok = tlexer.Peek(1)
        if ptok.TokType == lexer.TOK_EOF {
            instructionParsed = true
            break
        }
        
        core = eparser.getCore()
        ptok = tlexer.Peek(1)

        if ptok.TokType == lexer.TOK_SEMI_COLON {
             if ! PcFound {
                tlexer.ConsumeTillNewLine()
                continue
             }else{
				// Needs to be implemented	
			 }
             }
         }

        if ptok.TokType == lexer.TOK_ID || ptok.TokType == lexer.TOK_XREG || ptok.TokType == lexer.TOK_CSR {
             if ! PcFound {
               eparser.handle_reset_instr(core, instr)
               continue
            }   
        }

        if ptok.TokType == lexer.TOK_L_SQ_BRACKET {
            xtok = tlexer.GetToken() // This should consume the '['
            xtok = tlexer.GetToken() // This should have the PC
            instr.LineNum = uint64(xtok.LineNum)
            instr.PC, err = strconv.ParseUint(string(xtok.Value[2:]), 16, 64)
            if err != nil {
                fmt.Printf("Unable to convert value to uint64 : %s \n", xtok.Value)
                os.Exit(1)
            }
            instr.core = core
		
	    var nptok = tlexer.Peek(1)
	    if nptok.TokType == lexer.TOK_AT {
			instLineSeen = true
			xtok = tlexer.GetToken() // This should consume the '@' symbol
            eparser.handle_instruction_lines(core, instr)
	    	}
        }

}

func (eparser *EbbaParser) ebba_inst(instr *inst.Inst){
	var xtok *lexer.Token
	var ptok *lexer.Token
	var tlexer = eparser.tlexer

	ptok = tlexer.Peek(1)

	switch ptok.TokType {
	case lexer.TOJ_RV_JAL_TYPE:
		{
			xtok = tlexer.GetToken()
			instr.Inst = &xtok.Value

			if string(xtok.Value) == "jal"{
				xtok = tlexer.Match(lexer.TOK_XREG, "Expected an x register\n")
				xtok = tlexer.Match(lexer.TOK_COMMA, "Expected a comma\n")
			}
		}
	}
}


