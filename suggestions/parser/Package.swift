// swift-tools-version:5.3

import Foundation
import PackageDescription

var sources = ["src/parser.c"]
if FileManager.default.fileExists(atPath: "src/scanner.c") {
    sources.append("src/scanner.c")
}

let package = Package(
    name: "TreeSitterCodesuggestion",
    products: [
        .library(name: "TreeSitterCodesuggestion", targets: ["TreeSitterCodesuggestion"]),
    ],
    dependencies: [
        .package(name: "SwiftTreeSitter", url: "https://github.com/tree-sitter/swift-tree-sitter", from: "0.9.0"),
    ],
    targets: [
        .target(
            name: "TreeSitterCodesuggestion",
            dependencies: [],
            path: ".",
            sources: sources,
            resources: [
                .copy("queries")
            ],
            publicHeadersPath: "bindings/swift",
            cSettings: [.headerSearchPath("src")]
        ),
        .testTarget(
            name: "TreeSitterCodesuggestionTests",
            dependencies: [
                "SwiftTreeSitter",
                "TreeSitterCodesuggestion",
            ],
            path: "bindings/swift/TreeSitterCodesuggestionTests"
        )
    ],
    cLanguageStandard: .c11
)
