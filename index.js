// dom nodes marked with $ symbol
const elselect = ($parentElement, options) => {

    let $container;
    let $selectedValue;
    let $list;
    let $optionElements;
    let selectedOption;
    const placeholder = 'Please select';

    const setupDomElement = () => {
        $container = document.createElement('div');
        $container.classList.add('el-select');

        $selectedValue = document.createElement('div');
        $selectedValue.classList.add('selected-value');


        $list= document.createElement('div');
        $list.classList.add('list');

        $optionElements = options.availableOptions.map(availableOption => {
            const $element = document.createElement('div');
            $element.classList.add('list-option');
            $element.innerText = availableOption.name;
            $element.dataset.id = availableOption.id
            return $element;
        });


        const $emptyListOption = document.createElement('div');
        $emptyListOption.innerText = placeholder;
        $emptyListOption.classList.add('list-option');

        $optionElements = [$emptyListOption, ...$optionElements];

        $parentElement.append($container);
        $container.append($selectedValue);
        $container.append($list);
        $list.append(...$optionElements); // to append array use spread operator


    }
    const initializeListeners = () => {
        $selectedValue.addEventListener('click', () => {
            $list.classList.toggle('is-visible');
        });

        $optionElements.forEach($optionElement => {
            $optionElement.addEventListener('click', (e) => {
                $list.classList.remove('is-visible');
                setSelectedOption($optionElement.dataset.id)
            })
        });
    }

    const setSelectedOption = (selectedOptionId) => {
        if (!selectedOptionId) {
            selectedOption = null;
            $selectedValue.innerText = placeholder;
            return;
        }

        selectedOption = options.availableOptions.find(
            (availableOption) => availableOption.id === selectedOptionId
        );
        $selectedValue.innerText = selectedOption.name;
    }

    setupDomElement();
    setSelectedOption(options.selectedOptionId)
    initializeListeners();

};



const selectContainer = document.querySelector(".select-container");
languages = [
    {
        id: '1',
        name: "Javascript"
    },
    {
        id: '2',
        name: "Typescript"
    },
    {
        id: '3',
        name: "Node"
    },
    {
        id: '4',
        name: "Ruby"
    },
    {
        id: '5',
        name: "Go"
    },
    {
        id: '6',
        name: "Go"
    },
    {
        id: '7',
        name: "Go"
    },
];


elselect(selectContainer, {
    availableOptions: languages,
    selectedOptionId: "1" // default option selected
});