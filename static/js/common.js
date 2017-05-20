/**
 * Created by Mark on 5/7/17.
 */

function getToday() {
    var today = new Date();
    var thisMonth = today.getMonth()+1;
    return today.getFullYear() + "/" + zeroize(thisMonth) + "/" + today.getDate()
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

};

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

function getCookie() {
    var xsrf = $.cookie('_xsrf');
    var xsrflist = xsrf.split("|");
    var _xsrf = $.base64.decode(xsrflist[0])
    return _xsrf ? _xsrf : undefined;
}