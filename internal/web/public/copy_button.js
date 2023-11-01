class CopyButton extends HTMLButtonElement {
    connectedCallback() {
        this.innerText = 'Copy'

        this.linkedInputId = this.getAttribute('for');
        this.linkedInput = document.getElementById(this.linkedInputId);

        this.addEventListener('click', () => {
            if (!this.linkedInput) {
                return;
            }

            this.linkedInput.select();
            document.execCommand('copy');
            window.getSelection().removeAllRanges();

            this.innerText = 'Copied!';

            setTimeout(() => {
                this.innerText = 'Copy';
            }, 3000);
        });
    }
}

customElements.define('copy-button', CopyButton, {extends: 'button'});
