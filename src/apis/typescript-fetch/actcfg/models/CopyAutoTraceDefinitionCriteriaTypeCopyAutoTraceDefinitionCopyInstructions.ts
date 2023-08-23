/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity Management API
 * APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Instructions which will be considered when copying from the template. If this element is not sent, all the flags will be ignored.
 * @export
 * @interface CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions
 */
export interface CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions {
    /**
     * When true, this will copy the owner assignments to the trace definition.
     * @type {boolean}
     * @memberof CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions
     */
    allowOwnerAssignmentsCopy?: boolean;
}

/**
 * Check if a given object implements the CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions interface.
 */
export function instanceOfCopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructionsFromJSON(json: any): CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions {
    return CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructionsFromJSONTyped(json, false);
}

export function CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'allowOwnerAssignmentsCopy': !exists(json, 'allowOwnerAssignmentsCopy') ? undefined : json['allowOwnerAssignmentsCopy'],
    };
}

export function CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructionsToJSON(value?: CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'allowOwnerAssignmentsCopy': value.allowOwnerAssignmentsCopy,
    };
}

