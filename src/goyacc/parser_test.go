package goyacc

import "testing"

//Program structure

func TestParseEmptyProgramWithVars(t *testing.T) {
	input := `
		program testRun : var x, y: int; z, f: float; {
		
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseProgramNoVars(t *testing.T) {
	input := `
		program testRun : {
			if (x > 100.99) {} else {};
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

// Conditionals
func TestParseEmptyIf(t *testing.T) {
	input := `
		program testRun : var x, y: int; z, f: float; {
			if (x > 100.99) {};
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseEmptyIfElse(t *testing.T) {
	input := `
		program testRun : var x, y: int; z, f: float; {
			if (x > 100.99) {} else {};
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseNotEmptyIfElse(t *testing.T) {
	input := `
		program testRun : var x, y: int; z, f: float; {
			if (x > 100.99) {
				x = 10;
			} else {
				y = 20;
			};
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseNestedConditionals(t *testing.T) {
	input := `
		program testRun : var x, y: int; z, f: float; {
			if (x > 100.99) {
				x = 10;
				if (x < 0) {};
			} else {
				y = 20;
			    if (x < 0) {} else {};
			};
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

//Integration Test
func TestParseComplex(t *testing.T) {
	input := `
		program testRun : var x, y: int; z, f: float; {
			x = 10;
			y = 11;
			z = 100.2;
			f = z + y + x;
			m = z * x / (x-y);

			if (x+10.35 > 100) {
				print(x+10);
			};

			if (10.0 > 10) {
				print("this is illegal");
			} else {
				print("this is legal");
			};

			print("hi", "my","name is", "pacosw1");
			print(z > y);
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseValidNumericalAssignment(t *testing.T) {
	input := `
		program testRun : {
			x = 10;
			y = 10.10;
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseValid(t *testing.T) {
	input := `
		program testRun : {
			x = 10;
			y = 10.10;
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseValidPrintStatements(t *testing.T) {
	input := `
		program testRun : {
			print(x);
			print(x > 10);
			print(x,y,"hello world");
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestParseValidExpressionAssignment(t *testing.T) {
	input := `
		program testRun : {
			x = 10 + x / n * 1;
			y = 10;
		}
	`
	_, err := Parse(input)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestThrowError(t *testing.T) {
	input := `
		program testRun : {
			x = 10 +=- / n * 1;
			y = 10
		}
	`
	_, err := Parse(input)

	if err != nil {
		t.Logf(err.Error())

	}

	if err == nil {
		t.Fatalf("should not compile")
	}
}
