package templates

const (
	IndexTemplate = `<h1>LR 7</h1><p><a href="/quadratic-equation/">1) Quadratic Equation</a></p> <p><a href="/slice/">2) Slices</a></p>`

	SliceTemplate = `<h3>Working with Slices</h3>
    <div style="display: flex;">
        <div style="flex: 1; margin-right: 10px;">
            <p style="font-weight: bold;">Get Request</p>
            <form action="/slice/" method="GET" style="border: 1px solid #ccc; padding: 10px;">
                <p><label for="slice">Enter slice elements, using space as separator: </label>
                <input placeholder="3.1 0 5 6 0" name="slice" id="slice" style="margin-bottom: 5px; padding: 8px;"></p>
                <input type="submit" value="Calculate" style="padding: 10px; background-color: #0066cc; color: #fff; border: none; cursor: pointer;">
            </form>
        </div>
        <div style="flex: 1; margin-left: 10px;">
            <p style="font-weight: bold;">Post Request</p>
            <form action="/slice/" method="POST" style="border: 1px solid #ccc; padding: 10px;">
                <p><label for="slice">Enter slice elements, using space as separator: </label>
                <input placeholder="7.8 7 5.2 0 2 0" name="slice" id="slice" style="margin-bottom: 5px; padding: 8px;"></p>
                <input type="submit" value="Calculate" style="padding: 10px; background-color: #0066cc; color: #fff; border: none; cursor: pointer;">
            </form>
        </div>
    </div>`

	QuadraticEquationTemplate = `<h3 style="color: #333;">Solving Quadratic Equation</h3>
    <style>
        .form-container {
            display: flex;
            margin-bottom: 20px;
        }
        .form-container form {
            flex: 1;
            border: 1px solid #ccc;
            padding: 10px;
            margin-right: 10px;
        }
        label {
            font-weight: bold;
            margin-bottom: 5px;
            display: block;
        }
        input[type="submit"] {
            padding: 10px;
            background-color: #0066cc;
            color: #fff;
            border: none;
            cursor: pointer;
        }
        input[type="text"] {
            margin-bottom: 5px;
            padding: 8px;
            display: block;
        }
    </style>
    <div class="form-container">
        <form action="/quadratic-equation/" method="GET">
			<h1>Get Request</h1>
            <p><label for="a">Enter coefficient a:</label>
            <input placeholder="7" name="a" id="a"></p>
            <p><label for="b">Enter coefficient b:</label>
            <input placeholder="3" name="b" id="b"></p>
            <p><label for="c">Enter coefficient c:</label>
            <input placeholder="-9" name="c" id="c"></p>
            <input type="submit" value="Calculate">
        </form>
        <form action="/quadratic-equation/" method="POST">
			<h1>Post Request</h1>
            <p><label for="a">Enter coefficient a:</label>
            <input placeholder="-4" name="a" id="a"></p>
            <p><label for="b">Enter coefficient b:</label>
            <input placeholder="2" name="b" id="b"></p>
            <p><label for="c">Enter coefficient c:</label>
            <input placeholder="3" name="c" id="c"></p>
            <input type="submit" value="Calculate">
        </form>
    </div>`

	PageStart = `<!DOCTYPE HTML><html><head><title>LR-7</title><style>.error{color:#FF0000;}</style></head>`
	PageEnd   = `</html>`

	ErrorTag = `<p class="error">%s</p>`

	BodyStart = `<body>`
	BodyEnd   = `</body>`
)
