from selenium import webdriver
driver = webdriver.Chrome(executable_path='C:/Program Files/Google/Chrome/Application/chrome.exe')
# "C:\Program Files\Google\Chrome\Application\chrome.exe"
driver.get("https://www.amazon.com")
#for c in cookiestr.keys():
#    driver.add_cookie({'name':c,'value':cookiestr[c]})

#driver.get("https://www.amazon.com")

cookie = [item["name"] + "=" + item["value"] for item in driver.get_cookies()]
cookiestr = ';'.join(item for item in cookie)
print(cookiestr)