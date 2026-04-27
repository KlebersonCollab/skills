# Handoff Templates for Multi-Agent Workflows

## Template 1: Architecture to Development Handoff

```markdown
# Handoff: Feature [Feature Name]

## Context
You are now the **Lead Developer** for this feature. The Solution Architect has completed the design phase and you are responsible for implementation.

## What Was Decided

### Architecture
- [Brief description of the architecture]
- [Key technical decisions]
- [Technology stack chosen]

### Requirements
- [Functional requirement 1]
- [Functional requirement 2]
- [Non-functional requirement 1]

### Constraints
- [Constraint 1]
- [Constraint 2]

## Where to Find Specs

- **Design Spec**: `.specs/features/[feature-name]/design-spec.md`
- **API Contracts**: `.specs/api/[feature-name]-endpoints.md`
- **Component Specs**: `.specs/components/[feature-name]-components.md`
- **Database Schema**: `.specs/database/[feature-name]-schema.md`

## What You Must Do

1. **Invoke the appropriate skill**: `[skill-name]`
2. **Follow the implementation plan** in the design spec
3. **Create/update** the following files:
   - Implementation code in `src/`
   - Tests in `tests/`
   - Documentation in `docs/`
4. **Update** `STATE.md` with your progress
5. **Create** a handoff for the QA/SRE phase when done

## Success Criteria

- [ ] All requirements implemented
- [ ] All tests passing
- [ ] Code reviewed and approved
- [ ] Documentation updated
- [ ] Ready for QA handoff

## Next Steps

1. Read the design spec
2. Set up your development environment
3. Start implementation following the plan
4. Update STATE.md with your progress

## Questions?

If you have questions about the design, refer to the Decision Log in the design spec or ask the Architect for clarification.
```

## Template 2: Development to QA Handoff

```markdown
# Handoff: Feature [Feature Name] - QA Phase

## Context
You are now the **QA/SRE** for this feature. The Lead Developer has completed implementation and you are responsible for quality assurance and observability.

## What Was Implemented

### Changes Made
- [List of files changed]
- [List of new features added]
- [List of bugs fixed]

### Testing Status
- Unit tests: [X/Y passing]
- Integration tests: [X/Y passing]
- E2E tests: [X/Y passing]

### Known Issues
- [Any known issues or limitations]

## Where to Find Code

- **Source Code**: `src/[feature-name]/`
- **Tests**: `tests/[feature-name]/`
- **Documentation**: `docs/[feature-name]/`
- **API Endpoints**: `.specs/api/[feature-name]-endpoints.md`

## What You Must Do

1. **Invoke the appropriate skills**: `clean-code-mentor`, `observability-expert`
2. **Review the code** for:
   - SOLID principles compliance
   - Code quality and maintainability
   - Security vulnerabilities
   - Performance issues
3. **Test the feature**:
   - Manual testing
   - Automated testing
   - Edge cases
4. **Set up observability**:
   - Logging
   - Metrics
   - Tracing
5. **Update** `STATE.md` with your findings
6. **Create** a final handoff when approved

## Success Criteria

- [ ] Code review completed
- [ ] All tests passing
- [ ] No critical issues found
- [ ] Observability configured
- [ ] Documentation updated
- [ ] Ready for production

## Next Steps

1. Review the implementation
2. Run the test suite
3. Perform manual testing
4. Set up monitoring
5. Create final report

## Questions?

If you have questions about the implementation, refer to the code comments or ask the Developer for clarification.
```

## Template 3: Emergency Handoff

```markdown
# Emergency Handoff: [Issue Name]

## Context
This is an **emergency handoff**. The previous agent encountered an issue and needs you to take over immediately.

## The Problem

[Description of the problem or error]

## What Was Being Done

- [What the previous agent was working on]
- [What was completed]
- [What was in progress]
- [What was not started]

## Current State

- **Branch**: [branch name]
- **Files Modified**: [list of files]
- **Tests Status**: [passing/failing]
- **Error Messages**: [any error messages]

## What You Need to Do

1. **Assess the situation** - Read the error logs and understand the problem
2. **Stabilize** - Fix any critical issues preventing progress
3. **Continue** - Complete the task that was in progress
4. **Document** - Update any relevant documentation

## Resources

- **Error Logs**: [path to error logs]
- **Relevant Code**: [path to relevant code]
- **Documentation**: [path to relevant docs]

## Success Criteria

- [ ] Issue resolved
- [ ] Tests passing
- [ ] Task completed
- [ ] Documentation updated

## Notes

[Any additional notes or context]
```

## Template 4: Feature Completion Handoff

```markdown
# Handoff: Feature [Feature Name] - Complete

## Context
The feature [Feature Name] has been **completed** and is ready for production deployment.

## Summary

### What Was Built
- [Brief description of what was built]
- [Key features implemented]
- [Technical highlights]

### Quality Metrics
- Code coverage: [X%]
- Test pass rate: [X%]
- Performance: [metrics]
- Security: [status]

### Documentation
- [ ] API documentation updated
- [ ] User documentation updated
- [ ] Developer documentation updated
- [ ] CHANGELOG.md updated

## Deployment Checklist

- [ ] Code reviewed and approved
- [ ] All tests passing
- [ ] No critical issues
- [ ] Performance validated
- [ ] Security scan passed
- [ ] Documentation complete
- [ ] Rollback plan ready

## Post-Deployment

- [ ] Monitor for issues
- [ ] Collect metrics
- [ ] Gather feedback
- [ ] Update documentation as needed

## Next Steps

1. Deploy to production
2. Monitor for issues
3. Collect user feedback
4. Plan next iteration

## Lessons Learned

[Any lessons learned during development]
```

## Template 5: Daily Standup Handoff

```markdown
# Daily Standup Handoff - [Date]

## Context
This is a **daily standup handoff** to synchronize progress across agents.

## What I Did Yesterday

- [Task 1 completed]
- [Task 2 completed]
- [Task 3 in progress]

## What I'm Doing Today

- [Task 1]
- [Task 2]
- [Task 3]

## Blockers

- [Any blockers or issues]
- [Help needed]

## What's Next

- [Next steps for each task]
- [Dependencies on other agents]

## State Update

- **Current Branch**: [branch name]
- **Files Modified**: [list of files]
- **Tests Status**: [passing/failing]

## Notes

[Any additional notes or context]
```

## Best Practices

1. **Be Clear and Concise** - Use bullet points and clear language
2. **Include All Context** - Don't assume the next agent knows everything
3. **Provide Links** - Link to relevant files and documentation
4. **Define Success** - Clearly state what success looks like
5. **Update State** - Always update STATE.md with your progress
6. **Be Honest** - Report issues and blockers honestly
7. **Document Decisions** - Record why decisions were made

## Handoff Protocol

1. **Before Handoff**:
   - Complete your current task
   - Update STATE.md
   - Create handoff document
   - Run tests

2. **During Handoff**:
   - Provide clear context
   - Explain what was done
   - Explain what needs to be done
   - Answer questions

3. **After Handoff**:
   - Monitor progress
   - Be available for questions
   - Update documentation as needed