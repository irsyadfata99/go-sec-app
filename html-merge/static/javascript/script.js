document.addEventListener("DOMContentLoaded", function () {
  // Highlight active section in navbar
  const sections = document.querySelectorAll(".section");
  const navLinks = document.querySelectorAll(".nav-link");

  // Function to update active nav links
  function updateActiveNav() {
    let current = "";

    sections.forEach((section) => {
      const sectionTop = section.offsetTop;
      const sectionHeight = section.clientHeight;

      if (pageYOffset >= sectionTop - 200) {
        current = section.getAttribute("id");
      }
    });

    navLinks.forEach((link) => {
      link.classList.remove("active");
      if (link.getAttribute("href").substring(1) === current) {
        link.classList.add("active");
      }
    });
  }

  // Call the function on page load
  updateActiveNav();

  // Also call it on scroll
  window.addEventListener("scroll", updateActiveNav);

  // Animation on scroll
  const animateElements = document.querySelectorAll(".animate__animated");

  // Function to handle animations
  function handleAnimations() {
    animateElements.forEach((element) => {
      const elementPosition = element.getBoundingClientRect().top;
      const windowHeight = window.innerHeight;

      // If element is in viewport
      if (elementPosition < windowHeight - 100) {
        if (!element.classList.contains("animate__animated--triggered")) {
          element.classList.add("animate__animated--triggered");

          // For elements with fadeInUp animation
          if (element.classList.contains("animate__fadeInUp")) {
            // Reset animation by briefly removing and adding class
            element.classList.remove("animate__fadeInUp");
            setTimeout(() => {
              element.classList.add("animate__fadeInUp");
            }, 50);
          }
        }
      }
    });
  }

  // Set initial animations with a slight delay to ensure they work
  setTimeout(() => {
    animateElements.forEach((element) => {
      const elementPosition = element.getBoundingClientRect().top;
      const windowHeight = window.innerHeight;

      // If element is already in viewport on page load
      if (elementPosition < windowHeight) {
        // Make sure the animation is applied
        if (element.classList.contains("animate__fadeInUp")) {
          element.style.opacity = "1"; // Ensure element is visible
        }
      }
    });

    // Run initial animation check
    handleAnimations();
  }, 100);

  // Listen for scroll events to trigger animations
  window.addEventListener("scroll", handleAnimations);
});
