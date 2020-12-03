import os
import sys

import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import matplotlib.animation as ani


EPOCHS = None
RES_DIR = '../res'


def abort():
    print('Usage: python3 plot_results.py <EPOCHS>')
    exit(1)


if len(sys.argv) != 2:
    abort()

try:
    EPOCHS = int(sys.argv[1])
except:
    abort()

if EPOCHS < 1:
    abort()

fig = plt.figure()
plt.axis('off')

# read and store data from output files
frames = []
for i in range(EPOCHS):
    filename = RES_DIR + f'/rbn_{i}.csv' 
    data = pd.read_csv(filename, header=None, index_col=False).to_numpy()
    frames.append([plt.imshow(data, animated=True)])

# clean res directory
if os.name == 'nt':
    RES_DIR = RES_DIR.replace('/', '\\')
    os.system(f'del /Q {RES_DIR}')
else:
    os.system(f'rm -f {RES_DIR}/*')

# save the animation
animation = ani.ArtistAnimation(fig, frames, interval=250, blit=True, repeat_delay=1000)
animation.save(f'{RES_DIR}/RBN.gif')
