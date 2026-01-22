import React from 'react'
import EvidencePostForm from '../components/EvidencePostForm';
import EvidenceList from '../components/EvidenceList'

function Evidence() {
    return (
        <div>
            <h1>Evidence</h1>
            <EvidenceList/>
            <EvidencePostForm />
        </div>
    )
}

export default Evidence