
// var strDomian = "//www.iconfigurators.com/"
// var userid = 45958889
// var intConfigID = 17850;
// var intColor = 0;
// var strBody = ''
// var checkSelector = null;
// var bitPop = false;
// var title = "";
// var bitNoWheels = true;
// var vehicle = 0
// var part = '';
// var bodyStyle = '';
// var color =0
// var bid = 0
// $(function() {
// ICAPP.initChange();
// ICAPP.getRefVehicle();
// });
// var ICAPP= function ICAPP(){
// window.$ = jQuery;
// var getRefVehicle = function getRefVehicle(){
// title = $('#ic-vehicle-wrapper').attr('title');
// vehicle = $('#ic-vehicle-wrapper').attr('data-vehicleid');
// part = $('#ic-vehicle-wrapper').attr('data-part')
// if( part.length == 0 )
// nWheel = true;
// else
// nWheel = false
// getJSON(0,title,vehicle,part,nWheel);
// }
// var checkVehicle = function checkVehicle(data){
// var strResults = "";
// var strSelect = "";
// //var color = '';
// var body = '';
// // strResults = '<div id="vehicle-heading">';
// // strResults += title;
// // strResults += '</div>';
// strResults += '<div id="image-wrapper">';
// strResults += '<img id="vehicle-image" src="' + data.img[0].src +'" />';
// strResults += '</div>'
// strResults += '<div id="vehicle-selects">'
// if(data.Result > 0){
// // Create color select
// arrCID = data.img[0].colorID;
// arrName = data.img[0].colorName;
// arrImage = data.img[0].colorImage;
// // alert(data.wheel);
// if(data.wheel > 0)
// strColorSelect = '<ul id="vehicle-colors" >';
// else
// strColorSelect = '<ul id="vehicle-ref-colors" >';
// for(i=0; i < arrCID.length; i++){
// strColorSelect +='<li data-id="' + arrCID[i]+ '"><img src="//images.iconfigurators.com/images/vehicles/colors/swatch/' + arrImage[i] +'" /></li>';
// }
// strResults += strColorSelect;
// strResults += '</div>';
// //console.log(color)
// $('#ic-vehicle-wrapper').html(strResults);
// }
// if(data.wheel > 0){
// // Bind the color select
// $('#vehicle-colors li').bind('click',function() {
// color = $(this).attr('data-id');
// if(typeof option == 'undefined'){
// option = '';
// }
// getJSON(color,title,vehicle,part,false);
// });
// }else{
// $('#vehicle-ref-colors li').bind('click',function() {
// color = $(this).attr('data-id');
// // body = $('#vehicle-body').val();
// if(typeof option == 'undefined'){
// option = '';
// }
// getJSON(color,title,vehicle,'',true);
// });
// }
// // Bind the body select
// }
// var getJSON = function getJSON(color,title,vehicle, part, bitNoWheels){
// intColor = color;
// bitPop = true;
// // try{
// if(bitNoWheels)
// var url = strDomian + 'ap-json/get-AR-reference-image-color.aspx?vehicle=' +vehicle+'&color=' + color+ '&uid=' + userid ;
// else
// var url = strDomian + 'ap-json/ap-image-AR-part-id.aspx?vehicle=' +vehicle+'&part='+part+'&color=' + color + '& ID=' + intConfigID + '&uid=' + userid;
// $.ajax({
// dataType: 'jsonp',
// url: url,
// success: function(data){
// eval(data);
// }
// });
// /* }
// catch (err){
// //window.location.reload();
// }
// */
// }
// var initChange = function initChange(){
// //checkSelector = null;
// bitPop = false;
// // Add theme file to head
// $('head').append('<link rel="stylesheet" type="text/css" href="' + strDomian + 'pop/themes/aries.css" media="screen" />');
// // Hide button on load if no results
// $('.pop_up_vehicle').click(function(event) {
// event.preventDefault();
// var $this = $(this);
// title = $this.attr('title');
// remove = $this.attr('data-remove');
// var strCurPart = $this.attr('data-part');
// addRemovePart(remove,strCurPart);
// });
// }
// var replacePart = function replacePart(strNewPart,strOgPart){
// part = part.replace(strOgPart,strNewPart) ;
// getJSON(intColor,title,vehicle,part,false);
// }
// var addRemovePart = function addRemovePart(remove,strCurPart){
// var update = false;
// if( part.indexOf(strCurPart) > -1 && remove == 1 ){
// part = part.replace(strCurPart, "") ;
// update = true;
// }else if(part.indexOf(strCurPart) < 0 && remove == 0){
// update = true;
// if (part =='' )
// part = strCurPart;
// else
// part = part + ',' + strCurPart;
// }
// // console.log('remove: ' + remove + ' currentpart: ' + strCurPart + ' Index: ' + part.indexOf(strCurPart) + ' list: ' + part + ' update: ' + update )
// if(update){
// part = part.replace(",,", ",") ;
// if( part.indexOf(",") == 0 )
// part = part.substr(1, part.length);
// if( part.indexOf(",") == part.length ){
// part = part.substr(0, (part.length-1));
// }
// getJSON(intColor,title,vehicle,part,false);
// }
// }
// return {
// initChange: initChange,
// checkVehicle: checkVehicle,
// getRefVehicle:	getRefVehicle,
// getJSON: getJSON,
// replacePart:	replacePart,
// addRemovePart: addRemovePart
// }
// }();
// VehicleObj_ICA = function VehicleObj_ICA(JSONData){ICAPP.checkVehicle(JSONData);}
// // www.iconfigurators.com