class CopyButton extends HTMLButtonElement {
    connectedCallback() {
        this.linkedInputId = this.getAttribute('for');
        this.linkedInput = document.getElementById(this.linkedInputId);

        this.addEventListener('click', (e) => {
            if (!this.linkedInput) {
                return;
            }

            this.linkedInput.select();
            document.execCommand('copy');
            window.getSelection().removeAllRanges();
        });
    }
}

customElements.define('copy-button', CopyButton, {extends: 'button'});
