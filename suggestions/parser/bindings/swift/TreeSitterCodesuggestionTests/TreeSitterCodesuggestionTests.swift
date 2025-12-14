import XCTest
import SwiftTreeSitter
import TreeSitterCodesuggestion

final class TreeSitterCodesuggestionTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_codesuggestion())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Codesuggestion grammar")
    }
}
