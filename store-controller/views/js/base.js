"use strict";


// アロー関数出の書き方
const toggleMenu = () => {
  var $menu = document.querySelector('.header-menu');
  if ($menu.classList.contains('-show')) {
    $menu.classList.remove('-show');
  }
  else {
    $menu.classList.add('-show');
  }
}

window.addEventListener('load', function () {
  var $button = document.querySelector('.toggle-menu');

  // イベントの登録
  $button.addEventListener('click', toggleMenu);
});
