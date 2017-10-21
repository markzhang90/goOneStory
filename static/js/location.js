/**
 * Created by Mark on 6/11/17.
 */


function getLocation() {
    if (navigator.geolocation) {
        return navigator.geolocation.getCurrentPosition(showPosition, showError);
    } else {
        console.log("浏览器不支持地理定位。");
    }
    return null;

}

function showError(error) {
    console.log(222);
    switch (error.code) {
        case error.PERMISSION_DENIED:
            console.log("定位失败,用户拒绝请求地理定位");
            break;
        case error.POSITION_UNAVAILABLE:
            console.log("定位失败,位置信息是不可用");
            break;
        case error.TIMEOUT:
            console.log("定位失败,请求获取用户位置超时");
            break;
        case error.UNKNOWN_ERROR:
            console.log("定位失败,定位系统失效");
            break;
        default:
            console.log("定位失败");
            break;
    }
    return null;
}

// function showPosition(position){
//     var lat = position.coords.latitude; //纬度
//     var lag = position.coords.longitude; //经度
//     alert('纬度:'+lat+',经度:'+lag);
// }

function showPosition(position, callbackFunc) {

    var latlon = position.coords.latitude + ',' + position.coords.longitude;
    console.log(latlon);
    var data = null;
    //baidu
    var url = "https://api.map.baidu.com/geocoder/v2/?ak=C93b5178d7a8ebdb830b9b557abce78b&callback=renderReverse&location=" + latlon + "&output=json&pois=0";
    $.ajax({
        type: "GET",
        dataType: "jsonp",
        url: url,
        async: false,
        beforeSend: function () {
            console.log('正在定位...');
        },
        success: function (json) {
            if (json.status == 0) {
                data = json.result;
                var city = setLocationData(data);
                console.log(city);
                if(callbackFunc && city){
                    callbackFunc(city);
                }
                return data;
            }
        },
        error: function () {
            console.log(latlon + "地址位置获取失败");
        }
    });

    return data;
}


function getCityNameByLocation(callbackFunc) {
    var tryCity = getCookie("city");
    if (tryCity != null) {
        callbackFunc(tryCity);
        return tryCity;
    }
    var tryLoclat = getCookie("loclat");
    var tryLoclng = getCookie("loclng");
    // tryLoclat =39.978826;
    // tryLoclng = 116.4096263;
    if (tryLoclng != null && tryLoclat != null) {
        var position = {};
        position["coords"] = {};
        position["coords"]["latitude"] = tryLoclat;
        position["coords"]["longitude"] = tryLoclng;
        setCookie("loclat", tryLoclat, 2);
        setCookie("loclng", tryLoclng, 2);
        var allRes = showPosition(position, callbackFunc);
        return allRes;
    }
    var location = getLocation();
    return location;
}

function setLocationData(allRes) {

    if (allRes == null) {
        return null;
    }
    var addressComponent = allRes.addressComponent;
    console.log(addressComponent);
    var city = addressComponent.city;
    if (city != null) {
        setCookie("city", city, 2);
    }
    return city;
}
