var ImagePicker = Vue.extend({
    delimiters: ['${', '}'],
    template: '<div>' +
    '<div id="set-avatar" @mouseover="openDimmer" @mouseleave="closeDimmer" class="blurring dimmable image">' +
    '    <input id="uploader-1" type="file" style="display:none" @change="onFileChange">' +
    '    <div class="ui inverted dimmer">' +
    '          <div class="content">' +
    '<div class="center">' +
    '    <div class="ui primary button" @click="showModal">编辑头像</div>' +
    '</div>' +
    '          </div>' +
    '    </div>' +
    '    <img class="ui centered small circular image" :src="img_file">' +
    '</div>' +
    '<div id="avatar-picker" class="ui fullscreen modal">' +
    '    <div class="header">' +
    '    设置头像' +
    '    </div>' +
    '    <div class="content" style="min-height: 300px">' +
    '        <div class="ui grid" style="min-height: 300px">' +
    '            <div class="eight wide column">' +
    ' <div class="ui aligned center" style="height: 100%; width: 100%;">' +
    '     <img id="image-clip" :src="img_cache" alt="">' +
    ' </div>' +
    '            </div>' +
    '            <div class="eight wide column">' +
    ' <div class="ui small centered card">' +
    '     <div class="content">' +
    '         <div class="left floated author">' +
    '             <img class="ui avatar image" :src="img_file_test">' +
    '         </div>' +
    '     </div>' +
    '     <div class="image">' +
    '         <img :src="img_file_test">' +
    '     </div>' +
    '     <div @click="getImg" class="ui bottom attached button">' +
    '         <i class="add icon"></i>' +
    '             查看' +
    '     </div>' +
    ' </div>' +
    '            </div>' +
    '            <div class="ui header"></div>' +
    '            <p></p>' +
    '        </div>' +
    '    </div>' +
    '    <div class="actions">' +
    '        <div class="ui negative button">' +
    '            取消' +
    '        </div>' +
    '        <div @click="saveImg" class="ui green right labeled icon button"> <div :class="isActive ? active : disabled" class="ui inline small loader"></div>' +
    '            选好了，保存' +
    '        <i class="checkmark icon"></i>' +
    '    </div>' +
    '</div>' +
    '</div>' +
    '       </div>',
    props: [
        'tar_obj',
    ],
    data: function () {
        return {
            img_file: this.tar_obj.img,
            isActive: false,
            active: "active",
            disabled: "disabled",
            img_cache: this.tar_obj.img,
            cropper: null,
            img_file_test: this.tar_obj.img,
        }
    },
    methods: {

        openDimmer: function () {
            $('#set-avatar').dimmer('show');
        },
        closeDimmer: function () {
            $('#set-avatar').dimmer('hide');
        },
        openUploadDimmer: function () {
            $('#upload-avatar').dimmer('show');
        },
        closeUploadDimmer: function () {
            $('#upload-avatar').dimmer('hide');
        },
        uploadImage: function () {
            var _self = this;
            $('#uploader').click();
        },
        showModal: function () {
            $('#uploader-1').click();
        },
        onFileChange: function (e) {
            var _self = this;
            var files = e.target.files || e.dataTransfer.files;
            if (!files.length)
                return;
            var isImg = checkImgFile(files[0].type);
            if (!isImg)
                return;
            var reader = new FileReader();
            reader.readAsDataURL(files[0]);
            reader.onprogress = function (e) {
                reader.onload = function (e) {
                    // this 对象为reader
                    // reader.result 表示图片地址
                    _self.img_cache = reader.result;
                    if (_self.cropper != null) {
                        _self.cropper.replace(_self.img_cache);
                    } else {
                        var elem = document.getElementById('image-clip');
                        _self.cropper = new Cropper(elem, {
                            dragMode: 'move',
                            aspectRatio: 1,
                            restore: false,
                            guides: false,
                            center: false,
                            highlight: false,
                            cropBoxMovable: false,
                            cropBoxResizable: false,
                            toggleDragModeOnDblclick: false,
                        });
                    }
                }
            }

            _self.openAvatarPicker();
            $('#uploader-1').val('');
        },
        createImage: function (imgData) {
            var _self = this;
            var imageBase64 = imgData.replace("data:image/png;base64,", "");
            var formdata = new FormData();
            formdata.append('myfile', imageBase64);
            formdata.append('type', 'base64');
            formdata.append('_xsrf', getXsrfCookie("_xsrf"));
            _self.isActive = true
            $.ajax({
                type: "POST",
                url: "/uploader",
                data: formdata,
                cache: false,
                contentType: false,
                processData: false,
                dataType: "json",
                success: function (data) {
                    if (data.ErrNo == 0) {
                        // _self.img_file = data.Data.url
                        _self.tar_obj.upload_key = data.Data.key
                        _self.tar_obj.img = data.Data.url
                        _self.img_file = data.Data.url
                        _self.closeAvatarPicker()
                    } else {
                        alert("上传失败")
                    }
                    console.log(data);
                },
                complete: function () {
                    _self.isActive = false
                }
            });


        },
        openAvatarPicker: function () {
            $('#avatar-picker').modal({
                closable: true,
            });
            $('#avatar-picker').modal('show');
        },

        closeAvatarPicker: function () {
            $('#avatar-picker').modal('hide');
        },
        saveImg: function () {
            var _self = this;
            _self.getImg();
            _self.createImage(_self.img_file_test)
        },
        getImg: function () {
            var _self = this;
            var res = _self.cropper.getCroppedCanvas({
                width: 200,
                height: 200,
                minWidth: 256,
                minHeight: 256,
                maxWidth: 4096,
                maxHeight: 4096,
                fillColor: '#fff',
                imageSmoothingEnabled: false,
                imageSmoothingQuality: 'high',
            }).toDataURL();
            _self.img_file_test = res;

        }

    },
    watch: {
        // 如果 `question` 发生改变，这个函数就会运行
        img_cache: function () {
        }
    },
    computed: {},
    mounted: function () {
        var _self = this;
        var elem = document.getElementById('image-clip');
        _self.cropper = new Cropper(elem, {
            dragMode: 'move',
            aspectRatio: 1,
            autoCropArea: 0.65,
            restore: false,
            guides: false,
            center: false,
            highlight: false,
            cropBoxMovable: false,
            cropBoxResizable: false,
            toggleDragModeOnDblclick: false,
        });
    }

});
