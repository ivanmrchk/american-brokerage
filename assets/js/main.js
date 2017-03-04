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
    $("#form").submit(function(event) {

        var name = $("#name")
        var title = $("#title")
        var companyAddress = $("#companyAddress")
        var city = $("#city")
        var state = $("#state option:selected")
        var zip = $("#zip")
        var connfirmEmail = $("#confirmEmail")
        var pickup = $("#pickup")
        var drop = $("#drop")
        var email = $("#email")
        var phone = $("#phone")
        var companyName = $("#companyName")
        var message = $("#message")
        var subject = $("#subject")
        var last = $("#last")
        var mc = $("#mc")
        var dispatchName = $("#dispatchName")
        var dispatchPhone = $("#dispatchPhone")
        var dispatchEmail = $("#dispatchEmail")
        var insuranceCertificate = $("#insuranceCertificate")
        var w9 = $("#w9")
        var carrierAuthority = $("#carrierAuthority")
        var references = $("#references")
        var carrierBrokerAgreement = $("#carrierBrokerAgreement")
        var numberOfTrucks = $("#numberOfTrucks")
        var numberOfTrailers = $("#numberOfTrailers")
        var numberOfDriverTeams = $("#numberOfDriverTeams")
        checkField(numberOfDriverTeams)
        checkField(numberOfTrailers)
        checkField(numberOfTrucks)
        checkField(carrierBrokerAgreement)
        checkField(references)
        checkField(carrierAuthority)
        checkField(w9)
        checkField(insuranceCertificate)
        checkField(dispatchName)
        checkField(dispatchPhone)
        checkField(dispatchEmail)
        checkField(name);
        //checkField(title);
        checkField(companyAddress);
        checkField(city)
        checkField(state)
        checkField(zip)
        checkField(connfirmEmail)
        checkField(pickup)
        checkField(drop)
        checkField(subject)
        //checkField(last)
        checkField(email);
        checkField(phone);
        checkField(companyName);
        checkField(message);
        checkField(mc)





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


//scroll Conact
$(document).scroll(function() {
    contactScroll();
});

function contactScroll() {
    var y = window.scrollY;
    if (y > 505) {
        $('.make-it-fixed').css({
            "position": "fixed",
            "top": "69px"
        })

    } else if (y < 505) {
        $('.make-it-fixed').css({
            "position": "static"
        })
    }
}
