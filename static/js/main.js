// Tabz - Confirmation Request Pages
// Plain vanilla JS - no frameworks

document.addEventListener('DOMContentLoaded', function () {
  // Handle confirm button click
  var confirmBtn = document.querySelector('.btn--confirm');
  if (confirmBtn) {
    confirmBtn.addEventListener('click', function (e) {
      e.preventDefault();
      window.location.href = '/accepted';
    });
  }

  // Handle decline button click
  var declineBtn = document.querySelector('.btn--decline');
  if (declineBtn) {
    declineBtn.addEventListener('click', function (e) {
      e.preventDefault();
      window.location.href = '/denied';
    });
  }
});
