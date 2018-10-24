package dinner

/********
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import requests
import time
import logging
import random
from email.header import Header
from email.mime.text import MIMEText
from multiprocessing import Process
import smtplib
import sys

def func_meal():
 # 调用格式 http://124.127.188.104:1001/login
	hourofday = time.strftime('%H', time.localtime(time.time()))

            #获取session cookie
	get_session_cookie_header = {''}
    data = [{"username":"刘俊榜","passwd":2}]
	try:
		r = requests.post(url='http://124.127.188.104:1001/login', headers=get_session_cookie_header, json = (data))    # 最基本的GET请求
		r.raise_for_status()
	except:
		logging.error("signin_%s: error when get session cookie" % "1")
		continue
	else:
		logging.info("signin_%s: get session cookie succeed" % "1")

    # 配置日志
    logging.basicConfig(level=logging.DEBUG,
                        format='%(asctime)s %(filename)s[line:%(lineno)d] ==> %(levelname)s %(message)s',
                        datefmt='%Y-%m-%d %H:%M:%S',
                        filename='~/sign_in.log',
                        filemode='a')

    # 日志内容同时输出到控制台
    console = logging.StreamHandler()
    console.setLevel(logging.INFO)
    formatter = logging.Formatter('%(name)-12s: %(levelname)-8s %(message)s')
    console.setFormatter(formatter)
    logging.getLogger('').addHandler(console)


    max_failed_times = 88
    sign_result = 0
    error_message = ''

    while max_failed_times > 0:
        try:
            max_failed_times  = max_failed_times - 1
            # 此代码每天早上8点晚上6点执行一次，首先判断是否为工作日(http://www.k780.com/api/life.workday)，然后判断时间是否正确，然后打卡
            date = time.strftime('%Y%m%d', time.localtime(time.time()))

            # 调用格式 http://124.127.188.104:1001/login
            hourofday = time.strftime('%H', time.localtime(time.time()))

            #获取session cookie
            get_session_cookie_header = {''}
            data = [{"username":"刘俊榜","passwd":2}]
            try:
                r = requests.post(url='http://124.127.188.104:1001/login', headers=get_session_cookie_header, json = (data))    # 最基本的GET请求
                r.raise_for_status()
            except:
                logging.error("signin_%s: error when get session cookie" % "1")
                continue
            else:
                logging.info("signin_%s: get session cookie succeed" % "1")


            JSESSIONID = r.cookies['JSESSIONID']

            get_yjtToken_header = {'Content-Type': 'text/html; charset=utf-8',
                                 'apiVer': '1',
                                 'clientVersion': 'android_10755',
                                 'Connection': 'Keep-Alive',
                                 'User-Agent': 'yjt-oa',
                                 'Cookie2': '$Version=1'
                                 }

            #获取个人tocken
            cookies = {'JSESSIONID':JSESSIONID}
            data = {"contentId":0,"custUniqueId":0,"custName":"中国电信股份有限公司云计算分公司", "custVCode":0,"iccid":iccid,"password":password,"phone":phone,"userId":userId}
            try:
                r = requests.post(url='http://124.127.188.104:1001/', headers=get_yjtToken_header, json = (data), cookies= cookies)    # 最基本的GET请求
                r.raise_for_status()
            except:
                logging.error("signin_%s: error when get token" % name)
                continue
            else:
                logging.info("signin_%s: get yjttoken  succeed" % name)
            yjtToken = r.json()['payload']['yjtToken']

            #签到的http头部
            signin_header = {'Content-Type': 'application/json; charset=utf-8',
                                 'apiVer': '1',
                                 'clientVersion': 'android_10755',
                                 'Connection': 'Keep-Alive',
                                 'User-Agent': 'yjt-oa',
                                 'Cookie2': '$Version=1'
                                 }

            #签到啦，首先指定GPS经纬度，在两个矩形区域内随机选择
            latitude_low = [40.0535]
            latitude_high = [40.0571]
            longitude_low = [116.2888]
            longitude_high = [116.2912]

            region = random.randint(0,0)
            latitude = '%.6f' % random.uniform(latitude_low[region], latitude_high[region])
            longitude = '%.6f' % random.uniform(longitude_low[region], longitude_high[region])

            #开始构造http请求
            cookies = {'JSESSIONID':JSESSIONID, 'yjtToken': yjtToken}
            positionData = str(latitude) + ',' + str(longitude)
            data = {"descColor":0,"iccId":iccid,"id":0,"positionData":positionData,"positionDescription":"中国北京市海淀区软件园西二路","resultColor":0,"signResult":0,"type":"VISIT","userId":userId}
            try:
                r = requests.post(url='https://www.yijitongoa.com/yjtoa/s/signins/attendances', headers=signin_header, json = data, cookies= cookies)    # 最基本的GET请求
                r.raise_for_status()
            except:
                logging.error("signin_%s: error when sign use token" % name)
                continue
        except:
            time.sleep(random.randint(0,10))
            logging.error("signin_%s: error when signin" % name)
            continue
        else:
            logging.info("signin_%s: signin succeed" % name)
            sign_result = 1
            break

      #发送邮件通知
    if sign_result == 1:
        _send_email(phone, '成功', to_addr)
    else:
        if max_failed_times < 1:
            _send_email(phone, '失败', to_addr)


if __name__ == '__main__':

    # 登记信息
    iccids = ['89860116221100025000', '89860316780107221704', '89860318740101593738']
    imeis = ['866696020667786', '99000986138037', '86456503796758']
    phones = ['15210616814', '17326933267', '15330236736']
    passwords = ["bdt298032", '123456', 'passdocker123']
    userIds = [48244941, 48279104, '50965573']
    to_addrs = ['bit_bdt@foxmail.com', '492260444@qq.com', '75475760@qq.com']
    names = ['bdt', 'wlm', 'ljb']

    #*******************************通过对应的值判断是否需要签到,*************************************
    need_sign = [1,1,1]


    for i in range(names.__len__()):
        if (need_sign[i] == 1):
            p = Process(target=func_sign, args=(iccids[i], imeis[i], phones[i], passwords[i], userIds[i], to_addrs[i], names[i],))
            p.start()
*****/
