import {enableCharts} from "./chart/charts.js";
import {
    enableInvoiceAbortInsertButtons,
    enableInvoiceDeleteButtons,
    enableInvoiceEditButtons,
    enableInvoicePayButtons
} from "./invoices/invoices-button.js";


const enableDropDown = () => {
    const dropdownElementList = document.querySelectorAll('.dropdown');
    dropdownElementList.forEach(element => new coreui.Dropdown(element))
}


const enableHeaderToggler = () => {
    const elements = document.querySelectorAll('.header-toggler, .sidebar-toggler')
    elements.forEach(element => element.addEventListener('click', () => {
        coreui.Sidebar.getInstance(document.querySelector('#sidebar')).toggle()
    }));
}


const updateValori = () => {
    const elementList = document.querySelectorAll('.total-button');
    elementList.forEach(element => element.addEventListener('click', () => {
        const selector = element.getAttribute("desc-target")
        const targetList = document.querySelectorAll(selector);
        targetList.forEach(target => target.innerHTML = element.innerHTML);
    }));
}

updateValori();
enableHeaderToggler();
enableDropDown();
enableCharts();

enableInvoiceDeleteButtons();
enableInvoiceEditButtons();
enableInvoicePayButtons();


document.addEventListener("htmx:afterSwap", _ => {
    enableInvoiceDeleteButtons();
    enableInvoiceEditButtons();
    enableInvoicePayButtons();
    enableInvoiceAbortInsertButtons();
});

