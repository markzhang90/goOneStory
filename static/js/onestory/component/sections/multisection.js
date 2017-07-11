var templateMulti = '';
var MultiSection = Vue.extend({
    delimiters : ['${', '}'],
  	template: '<div class="ui clearing segment">\
  					<div class="ui two column grid">\
  						<div class="column">\
		  					<div :id=getImgSectionId class="blurring dimmable image">\
		  						<input :id=getImgUploadId type="file" style="display:none" @change="onFileChange">\
		  						<div class="ui inverted dimmer">\
		  							<div class="content">\
		  								<div class="center">\
		  									<div class="ui black basic button" @click="uploadImage()">上传图片  <div :class="isActive ? active : disabled" class="ui inline small loader"></div></div>\
										</div>\
									</div>\
								</div>\
								<img class="ui image" @mouseover="openDimmer()" :src=img_file >\
							</div>\
						</div>\
						<div class="column">\
							<div class="field ui reply form">\
	  							<textarea v-on:input="updateValue()" v-model="para"></textarea>\
	  						</div>\
	  					</div>\
  					</div>\
  				</div>',
  				
  	props: [
    	'tar_obj',
    ],

    data: function(){
    	return {
    		para: this.tar_obj.para,
    		img_file: this.tar_obj.upload_img,
            isActive: false,
            active: "active",
            disabled: "disabled",
    	}
    },
    methods: {

    	openDimmer: function(){
            var _self = this;
			$('#'+_self.getImgSectionId).dimmer('show');
    	},
    	closeDimmer: function(){
            var _self = this;
			$('#'+_self.getImgSectionId).dimmer('hide');
    	},
    	uploadImage: function(){
            var _self = this;
    		$('#'+_self.getImgUploadId).click();
    	},
    	onFileChange: function(e) {
	      	var files = e.target.files || e.dataTransfer.files;
            var _self = this;
	      	if (!files.length)
	        	return;
	      	var _self = this;
			// _self.tar_obj.upload_file = files[0]
            _self.isActive = true
	      	this.createImage(files[0]);
	    },
        createImage: function(file) {
            var image = new Image();
            var _self = this;
            var formdata = new FormData();
            formdata.append('myfile', file);
            formdata.append('_xsrf', getXsrfCookie("_xsrf"));
            $.ajax({
                type: "POST",
                url:  "/uploader",
                data: formdata,
                cache : false,
                contentType : false,
                processData : false,
                dataType: "json",
                success: function (data) {
                    if (data.ErrNo == 0){
                        // _self.img_file = data.Data.url
                        loadImage(data.Data.url, _self.updateImage)
                        _self.tar_obj.upload_key = data.Data.key
                    }else{
                        alert("上传失败")
                    }
                    console.log(data);
                },
            });
        },

        updateImage:function (url) {
            var _self = this;
            _self.img_file = url
            _self.isActive = false
            _self.closeDimmer()
        },

	   	updateValue: function(){
    		var _self = this
        	this.tar_obj.para = _self.para
    	}
    },
    computed: {
    	getImgSectionId: function(){
            var _self = this;
    		return 'img_section_' + _self.tar_obj.finder;
    	},

    	getImgUploadId: function(){
            var _self = this;
    		return 'img_upload_' + _self.tar_obj.finder;
    	},
    }
});
