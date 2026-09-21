print("Hello, World!")
print("\n")
print("All required libraries have been imported successfully!")

# Core imports
import sys
import jupyter_core
import pandas as pd
import numpy as np
import json

# Optional image libraries with safe version check
try:
    import openai
    print(f"OpenAI SDK version: {openai.__version__}")
except ImportError:
    print("OpenAI SDK not installed.")

try:
    import diffusers
    print(f"Diffusers version: {diffusers.__version__}")
except ImportError:
    print("Diffusers not installed.")

# Environment verification
print("\nEnvironment check:")
print(f"Python version:         {sys.version}")
print(f"Jupyter Core version:   {jupyter_core.__version__}")
print(f"Pandas version:         {pd.__version__}")
print(f"Numpy version:          {np.__version__}")

# JSON test
data = {"city": "Lagos", "temperature": 31}
print("JSON module works correctly:", json.dumps(data))