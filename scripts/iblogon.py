from selenium import webdriver
from selenium.webdriver.chrome.options import Options
import os
import time 
import socket

# sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
# result = sock.connect_ex(('127.0.0.1',5000))
# if result == 0:
   # print("Port is open")
   # result.close()
# else:
   # print("Port is not open")
os.system('kill $(lsof -t -i:5000)')

startWebPortal = './bin/run.sh root/config.yml &'
os.system(startWebPortal)
time.sleep(4)

#driver = webdriver.Chrome()
# CHROME_PATH = '/usr/bin/google-chrome'
# CHROMEDRIVER_PATH = '/usr/bin/chromedriver'
# WINDOW_SIZE = "1920,1080"
# chrome_options = Options()  
# chrome_options.add_argument("--headless")  
# chrome_options.add_argument('--ignore-certificate-errors')
# # chrome_options.add_argument("--window-size=%s" % WINDOW_SIZE)
# # chrome_options.binary_location = CHROME_PATH
# chrome_options.add_argument('--no-sandbox') # # Bypass OS security model
# chrome_options.add_argument('start-maximized')
# chrome_options.add_argument('disable-infobars')
# chrome_options.add_argument("--disable-extensions")

print("Sending credential")
driver = webdriver.Chrome()  
driver.get("https://localhost:5000")
driver.find_element_by_id("user_name").send_keys("xxxx")
driver.find_element_by_id("password").send_keys("xxxx")
driver.find_element_by_id("submitForm").click()
# driver.close()

