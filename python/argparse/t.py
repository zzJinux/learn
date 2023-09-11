import argparse

parser = argparse.ArgumentParser()
parser.add_argument('pos', help='positional arg')
parser.add_argument('cal-asd', help='positional arg')
parser.add_argument('--zzz', action='store_true', help='ha-ha help')
parser.add_argument('--foo-bar',action='store_true',  help='foo help')

print(parser.parse_args(['--foo-bar', '--zzz' , 'AAA', 'BBB']))
