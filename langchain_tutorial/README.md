# Langchain Tutorial from DeepLearning.ai

June 1, 2023
https://www.deeplearning.ai/short-courses/langchain-for-llm-application-development/


Related courses:

How Diffusion Models work

https://learn.deeplearning.ai/diffusion-models/lesson/1/introduction

# Launching

$ jupyter-lab

The API key for the OpenAPI API is in the file called .env. It is read by a function
in 01-basic.ipynb and this notebook is executed by other notebooks.

There are pip-based installation of certain modules dispersed throughout the code,
but ideally they should all be placed in a single notebook for a one-time run for 
each kernel. (Each notebook has its own kernel).


# Other tips


Use "pip install" commands to install the necessary packages, and then start the
jupyter-lab tool. 

jupyter-lab
panel
dotenv
openai

Make sure you get your own key and place it in the directory where the jupyter-lab
server is running. Otherwise fix the code to use find_dotenv() to work properly.

The .env file should be of the format, use your own API-key here as the value of the parameter
Ensure that you do not print this key in your notebooks or commit the .env file as 
part of your code repository. 

OPENAI_API_KEY=sk-asdfklasd8fadsfadskfadsf

(check the spelling and the underscores..)

$ jupyter-lab --port 8889 
- will start another instance of the jupyter-lab server on a different port. The default
port is 8888.





