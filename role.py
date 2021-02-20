# coding: utf-8  
import os
import requests
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait as Wait
from selenium.webdriver.support import expected_conditions as Expect
from langconv import * 


header = {
    'Accept':'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
    'Accept-Encoding':'gzip',
    'Accept-Language':'zh-CN',
    'Cache-Control':'no-cache',
    # 'Connection':'keep-alive',
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3163.100 Safari/537.36'
}


def roleAvatarDownload():
    if os.path.exists(os.getcwd() +'/static') is False:  # mac路径
         os.mkdir(os.getcwd()+'/static')
         os.mkdir(os.getcwd()+'/static/N')
         os.mkdir(os.getcwd()+'/static/SR')
         os.mkdir(os.getcwd()+'/static/SSR')
         print('成功创建文件夹')
    print(os.getcwd())   
    browser = webdriver.Chrome()
    browser.get('https://pcredivewiki.tw/Character')
    Wait(browser, 50).until(Expect.presence_of_element_located((By.CSS_SELECTOR, ".card ")))
    r = requests.session()
    r.headers = header
    r.keep_alive = False # 关闭多余连接避免爬取错误
    elem = browser.find_elements_by_css_selector('.pb-3')
    for cardContainer in elem:
        cardType = cardContainer.find_element_by_css_selector('.item-title').text
        cardItem = cardContainer.find_elements_by_css_selector('.card')
        for card in cardItem:
            fileName = '' 
            img = card.find_element_by_css_selector('.img-fluid')
            link = img.get_attribute('src')
            name = card.find_element_by_css_selector('.text-muted').text
            # print(cardType)
            if cardType == "3星":
                filename = os.getcwd() + '/static/SSR/' + Converter('zh-hans').convert(name) + '.png'  
            elif cardType == "2星":
                filename = os.getcwd() + '/static/SR/' + Converter('zh-hans').convert(name)  + '.png'  
            else:
                filename = os.getcwd() + '/static/N/' + Converter('zh-hans').convert(name)  +  '.png'     
            if os.path.exists(filename) is False:
                with open(filename, 'wb') as f:
                    
                    print('正在下载图片 -  {0}'.format(filename))
                    f.write(r.get(link).content)
            else: 
                print('{0} 已下载'.format(filename))
    print('下载完毕！')        
    browser.close()


roleAvatarDownload()
