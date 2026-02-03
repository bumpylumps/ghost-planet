import React from 'react'
import TextEvidencePostForm from '../components/TextEvidencePostForm';
import EvidenceList from '../components/EvidenceList'

function Evidence() {
    return (
        <div>
            <h1>Evidence</h1>
            <EvidenceList/>
            <TextEvidencePostForm />
        </div>
    )
}

export default Evidence