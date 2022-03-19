from parser import parser

def runTest():

    print("Parsing code")
    f = open('correct.ld')
    s = f.read()
    f.close()

    parser.parse(s)

    print('Parsed Successfully')

if __name__ == "__main__":
    runTest()