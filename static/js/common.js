/**
 * Created by Mark on 5/7/17.
 */

function getToday() {
    var today = new Date();
    var thisMonth = today.getMonth()+1;
    return today.getFullYear() + "/" + zeroize(thisMonth) + "/" + zeroize(today.getDate())
}

function changeDate(dateToGo, pidkcedData, type){

    var b = arguments[2] ? arguments[2] : 1;

    var _self = this
    var converted = Date.parse(pidkcedData);
    var curDate = new Date(converted);
    if(type == 1){
        var newDate = new Date(curDate.getTime() + 24*60*60*1000*dateToGo);  //后一天
    }else{
        var newDate = new Date(curDate.getTime() - 24*60*60*1000*dateToGo);  //前一天
    }
    return stringToDate(newDate);

}

function zeroize(value, length) {

    if (!length) length = 2;

    value = String(value);

    for (var i = 0, zeros = ''; i < (length - value.length); i++) {

        zeros += '0';

    }

    return zeros + value;

}

function dateIntToData(date) {
    var this_date = date / 10000;
    var year = parseInt(this_date)
    var rest = date % 10000;
    var month = parseInt(rest / 100)
    var day = parseInt(rest % 100)
    return year + "/" + zeroize(month) + "/" + day
}

function stringToDate(DateStr){

    var converted = Date.parse(DateStr);

    var myDate = new Date(converted);
    if (isNaN(myDate)){
        var arys= DateStr.split('-');
        myDate = new Date(arys[0],arys[1],arys[2]);
    }
    month = myDate.getMonth() + 1;
    return myDate.getFullYear() + "/" + zeroize(month) + "/" + myDate.getDate();
}

function getXsrfCookie() {
    var xsrf = $.cookie('_xsrf');
    var xsrflist = xsrf.split("|");
    var _xsrf = $.base64.decode(xsrflist[0])
    return _xsrf ? _xsrf : undefined;
}

function setCookie(name,value, expire)
{
    if(expire == null){
        expire = 7
    }
    $.cookie(name, value, { expires: expire });
}

function redirectLogin() {
    passId = getCookie('passid')
    if(passId && passId!=null){
        return true;
    }
    setCookie("passid", "", -1);
    setCookie("city", "", -1);
    setCookie("loclat", "", -1);
    setCookie("loclng", "", -1);

    if (!passId){
        window.location.href="/user/logintosys"
    }
    return false;
}

function redirectLogout() {
    setCookie("passid", "", -1);
    window.location.href="/"
}

function getCookie(name)
{
    var result = $.cookie(name);
    if (result == null || result.length == 0){
        return null;
    }
    return result;
}

function loadImage(url, callback) {
    var img = new Image(); //创建一个Image对象，实现图片的预下载
    img.src = url;

    if(img.complete) { // 如果图片已经存在于浏览器缓存，直接调用回调函数
        callback(img.src);
        return; // 直接返回，不用再处理onload事件
    }
    img.onload = function () { //图片下载完毕时异步调用callback函数。
        callback(img.src);//将回调函数的this替换为Image对象
    };
}


function loadUserInfo (callbackSucc, callBackFail) {
    redirectLogin();
    $.ajax({
        type: "Get",
        url: "/user/getuserprofile",
        data: {},
        dataType: "json",
        success: function (data) {
            callbackSucc(data);
        },
        error: function (e) {
            callBackFail();
        },
    })
}

/**
 *
 * @param id
 * @param callbackFunc
 */
function loadPostData(id, callbackFunc) {

    if(isNaN(id)){
        callbackFunc()
        return;
    }
    var inputdata = {};
    inputdata['id'] = id;
    $.ajax({
        type: "GET",
        url: "/post/getuserpostbyid",
        data: inputdata,
        dataType: "json",

        success: function (data) {
            callbackFunc(data)
        },
        error: function () {
            callbackFunc()
        }
    });
}

/**
 *
 * @param splitDate
 * @param left
 * @param callbackFunc
 */
function loadRecordByDateRange(splitDate, left, callbackFunc){
    var requestData = {};
    if (left){
        requestData.start = splitDate;
        requestData.order = "asc";
    }else{
        requestData.end = splitDate;
        requestData.order = "desc";
    }
    requestData.limit = 6;
    requestData.total = 1;
    $.ajax({
        type: "GET",
        url: "/post/getuserpostdaterange",
        data: requestData,
        dataType: "json",
        success: function (data) {
            callbackFunc(data)
        },
        error: function () {
            callbackFunc()
        }
    });
}

/**
 * @param element
 */
function autoProgressRun(element) {
    clearInterval(window.fakeProgress)
    element.progress('reset');
    element.progress('remove success');
    // updates every 10ms until complete
    window.fakeProgress = setInterval(function() {
;        // stop incrementing when complete
        if(element.progress('is complete') || element.progress('get percent') == 90) {
            clearInterval(window.fakeProgress)
        }else{
            element.progress({
                percent: element.progress('get percent') + 10
            });
        }
    }, 300);
}

/**
 * @param element
 */
function closeProgressRun(element) {
    element.progress({
        percent: 100
    });
}

function Init(callback, callbackfail) {
    $.ajax({
        type: "GET",
        url: "/api/wechat/initinfo",
        data: {},
        dataType: "json",
        success: function (data) {
            callback(data)
        },
        error: function () {
            callbackfail()
        }
    });
}

function checkImgFile(fileType) {
    if(!/image\/\w+/.test(fileType)){
        alert("文件必须为图片！");
        return false;
    }
    return true;
}
