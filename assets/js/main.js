// makes sure the whole site is loaded
jQuery(window).load(function() {
    // will first fade out the loading animation
    jQuery("#status").fadeOut();
    // will fade out the whole DIV that covers the website.
    jQuery("#preloader").delay(1000).fadeOut("slow");
})

jQuery(document).ready(function($) {
    'use strict';

    function checkField(fieldName) {
        if (fieldName.val() == "") {
            fieldName.css({
                "border": "1px solid red"
            })
            event.preventDefault()
        }
        fieldName.on('click', function() {
            $(this).css({
                "border": "1px solid rgba(0, 0, 0, .2)"
            })
        })
    }
    // validator
    // get quote
    $("#getQuote").submit(function(event) {

        var qname = $("#qname")
        var qemail = $("#qemail")
        var qphone = $("#qphone")
        var qcompanyName = $("#qcompanyName")
        var qmessage = $("#qmessage")
        checkField(qname);
        checkField(qemail);
        checkField(qphone);
        checkField(qcompanyName);
        checkField(qmessage);



    })
    // contact modal
    $("#contactModal").submit(function(event) {

        var cname = $("#cname")
        var cemail = $("#cemail")
        var cphone = $("#cphone")
        var cmessage = $("#cmessage")
        checkField(cname);
        checkField(cemail);
        checkField(cphone);
        checkField(cmessage);

    })
    $("#contactFormPage").submit(function(event) {

        var cfname = $("#cfname")
        var cfemail = $("#cfemail")
        var cfphone = $("#cfphone")

        var cfmessage = $("#cfmessage")
        checkField(cfname);
        checkField(cfemail);
        checkField(cfphone);

        checkField(cfmessage);

    })

    // Search
    $('a[href="#search"]').on('click', function(event) {
        event.preventDefault();
        $('#search').addClass('open');
        $('#search > form > input[type="search"]').focus();
    });

    $('#search, #search button.close').on('click keyup', function(event) {
        if (event.target == this || event.target.className == 'close' || event.keyCode == 27) {
            $(this).removeClass('open');
        }
    });



    // Menu
    $('nav#mobile-menu').mmenu({
        extensions: ['effect-slide-menu', 'pageshadow'],
        searchfield: true,
        counters: true,
        navbar: {
            title: 'Menu'
        },
        navbars: [{
            position: 'top',
            content: ['searchfield']
        }, {
            position: 'top',
            content: [
                'prev',
                'title',
                'close'
            ]
        }, {
            position: 'bottom',
            content: []
        }]
    });

    // Owl-Carousel
    $('.testimonials-carousel').owlCarousel({
        loop: true,
        margin: 30,
        responsiveClass: true,
        responsive: {
            0: {
                items: 1,
                nav: false
            },
            600: {
                items: 1,
                nav: false
            },
            1000: {
                items: 2,
                nav: false,
                loop: false
            }
        }
    });

    $('.talking-carousel').owlCarousel({
        loop: true,
        margin: 30,
        dots: false,
        responsiveClass: true,
        responsive: {
            0: {
                items: 2,
                nav: false
            },
            600: {
                items: 1,
                nav: false
            },
            1000: {
                items: 1,
                nav: false,
                loop: false
            }
        }
    })

    // SLIP HOVER CLIENTS
    $(".our-clients").sliphover({
        backgroundColor: 'rgba(202,31,38,0.75)'
    });

    // Maps
    var gmMapDiv = $("#map-canvas");

    (
        function($) {

            if (gmMapDiv.length) {

                var gmMarkerAddress = gmMapDiv.attr("data-address");
                var gmHeight = gmMapDiv.attr("data-height");
                var gmWidth = gmMapDiv.attr("data-width");
                var gmZoomEnable = gmMapDiv.attr("data-zoom_enable");
                var gmZoom = gmMapDiv.attr("data-zoom");

                gmMapDiv.gmap3({
                    action: "init",
                    marker: {
                        address: gmMarkerAddress,
                        options: {
                            icon: "http://transport.thememove.com/wp-content/themes/tm_transport/images/map-marker.png",
                        },
                    },
                    map: {
                        options: {
                            zoom: parseInt(gmZoom),
                            zoomControl: true,
                            mapTypeId: google.maps.MapTypeId.ROADMAP,
                            mapTypeControl: false,
                            scaleControl: false,
                            scrollwheel: gmZoomEnable == 'enable' ? true : false,
                            streetViewControl: false,
                            draggable: true,
                            styles: [{
                                "featureType": "all",
                                "elementType": "all",
                                "stylers": [{
                                    "saturation": -100
                                }, {
                                    "gamma": 0.5
                                }]
                            }],
                        }
                    }
                }).width(gmWidth).height(gmHeight);
            }
        }
    )(jQuery);

    // lightSlider - Single product page
    $('#imageGallery').lightSlider({
        gallery: true,
        item: 1,
        thumbItem: 3,
        slideMargin: 0,
        enableDrag: false,
        currentPagerPosition: 'left',
        adaptiveHeight: true,
        onSliderLoad: function(el) {
            el.lightGallery({
                selector: '#imageGallery .lslide'
            });
        }
    });

    // Ship different
    $('.different-collapse').collapse();

    $('.ship-different .custom-heading a input[type=checkbox], .ship-different .custom-heading a label').on('click', function(e) {
        e.stopPropagation();
        $(this).parent().trigger('click');
    });

    $('#collapseDifferentAddress').on('show.bs.collapse', function(e) {
        if (!$('.ship-different .custom-heading a input[type=checkbox]').is(':checked')) {
            return false;
        }
    });

});
