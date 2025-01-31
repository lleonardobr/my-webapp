// Efeito de digitação no título
const heroTitle = document.querySelector('.hero h2');
const text = "Automate. Deploy. Scale.";
let index = 0;

function typeWriter() {
    if (index < text.length) {
        heroTitle.innerHTML += text.charAt(index);
        index++;
        setTimeout(typeWriter, 100);
    }
}

typeWriter();

// Efeito de hover nos cards de features
const features = document.querySelectorAll('.feature');
features.forEach(feature => {
    feature.addEventListener('mouseenter', () => {
        feature.style.transform = 'translateY(-10px)';
        feature.style.boxShadow = '0 0 20px rgba(0, 255, 204, 0.5)';
    });
    feature.addEventListener('mouseleave', () => {
        feature.style.transform = 'translateY(0)';
        feature.style.boxShadow = '0 0 10px rgba(0, 255, 204, 0.3)';
    });
});