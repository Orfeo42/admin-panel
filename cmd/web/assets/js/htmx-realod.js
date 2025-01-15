const buttonDeleteClick = (element) => {
    Swal.fire({
        title: 'Vuoi eliminare?',
        showCancelButton: true,
        confirmButtonText: 'Elimina',
        text: 'Vuoi procedere all\'eliminazione di questo elemento?',
    }).then((result) => {
        if (result.isConfirmed) {
            htmx.trigger(element, 'delete')
        }
    });
}

const enableDeleteButtons = () => {
    const elements = document.querySelectorAll('.invoice-delete-button')
    elements.forEach(element => element.addEventListener('click', () => buttonDeleteClick(element)));
}

const buttonEditClick = (element) => {
    let editing = document.querySelector('.editing')
    if (editing) {
        Swal.fire({
            title: 'Stai già modificando!',
            showCancelButton: true,
            confirmButtonText: 'Sì, Modifica questa riga!',
            text: 'Stai già modificando un altra riga. Vuoi annullare per modificare questa?'
        }).then((result) => {
            if (result.isConfirmed) {
                htmx.trigger(editing, 'cancel')
                htmx.trigger(element, 'edit')
            }
        })
        return;
    }
    htmx.trigger(element, 'edit')
}

const enableEditButtons = () => {
    const elements = document.querySelectorAll('.invoice-edit-button')
    elements.forEach(element => element.addEventListener('click', () => buttonEditClick(element)));
}

document.addEventListener("htmx:afterSwap", _ => {
    enableDeleteButtons();
    enableEditButtons();
});

const buttonPayClick = (element) => {
    Swal.fire({
        title: 'Vuoi impostare come pagata questa fattura?',
        showCancelButton: true,
        confirmButtonText: 'Paga',
        text: 'Vuoi procedere all\'impostare questa fattura pagata?',
    }).then((result) => {
        if (result.isConfirmed) {
            htmx.trigger(element, 'pay')
        }
    });
}

const enablePayButtons = () => {
    const elements = document.querySelectorAll('.invoice-pay-button')
    elements.forEach(element => element.addEventListener('click', () => buttonPayClick(element)));
}

document.addEventListener("htmx:afterSwap", _ => {
    enableDeleteButtons();
    enableEditButtons();
    enablePayButtons();
});

enableDeleteButtons();
enableEditButtons();
enablePayButtons();
