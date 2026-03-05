document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.querySelector('.login-form');

    loginForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        
        console.log('Login Attempt:', { email, password });
        alert('Login submitted! Check console for details.');
    });

    const socialBtns = document.querySelectorAll('.social-btn');
    socialBtns.forEach(btn => {
        btn.addEventListener('click', () => {
            const provider = btn.innerText.trim();
            console.log(`Social Login with: ${provider}`);
            alert(`Redirecting to ${provider} login...`);
        });
    });
});
