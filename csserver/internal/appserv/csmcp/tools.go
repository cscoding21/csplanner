package csmcp

import "github.com/mark3labs/mcp-go/mcp"

// Add a calculator tool
var calculatorTool = mcp.NewTool("calculate",
	mcp.WithDescription("Perform basic arithmetic operations"),
	mcp.WithString("operation",
		mcp.Required(),
		mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
		mcp.Enum("add", "subtract", "multiply", "divide"),
	),
	mcp.WithNumber("x",
		mcp.Required(),
		mcp.Description("First number"),
	),
	mcp.WithNumber("y",
		mcp.Required(),
		mcp.Description("Second number"),
	),
)
