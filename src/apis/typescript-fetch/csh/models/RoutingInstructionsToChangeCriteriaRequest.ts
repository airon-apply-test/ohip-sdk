/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CompRoutingRequestType } from './CompRoutingRequestType';
import {
    CompRoutingRequestTypeFromJSON,
    CompRoutingRequestTypeFromJSONTyped,
    CompRoutingRequestTypeToJSON,
} from './CompRoutingRequestType';
import type { RoutingInstructionType } from './RoutingInstructionType';
import {
    RoutingInstructionTypeFromJSON,
    RoutingInstructionTypeFromJSONTyped,
    RoutingInstructionTypeToJSON,
} from './RoutingInstructionType';

/**
 * Comp Accounting Request routing
 * @export
 * @interface RoutingInstructionsToChangeCriteriaRequest
 */
export interface RoutingInstructionsToChangeCriteriaRequest {
    /**
     * 
     * @type {CompRoutingRequestType}
     * @memberof RoutingInstructionsToChangeCriteriaRequest
     */
    compRequestInfo?: CompRoutingRequestType;
    /**
     * Set of routing instructions associated to this routing type.
     * @type {Array<RoutingInstructionType>}
     * @memberof RoutingInstructionsToChangeCriteriaRequest
     */
    instructions?: Array<RoutingInstructionType>;
}

/**
 * Check if a given object implements the RoutingInstructionsToChangeCriteriaRequest interface.
 */
export function instanceOfRoutingInstructionsToChangeCriteriaRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoutingInstructionsToChangeCriteriaRequestFromJSON(json: any): RoutingInstructionsToChangeCriteriaRequest {
    return RoutingInstructionsToChangeCriteriaRequestFromJSONTyped(json, false);
}

export function RoutingInstructionsToChangeCriteriaRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingInstructionsToChangeCriteriaRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'compRequestInfo': !exists(json, 'compRequestInfo') ? undefined : CompRoutingRequestTypeFromJSON(json['compRequestInfo']),
        'instructions': !exists(json, 'instructions') ? undefined : ((json['instructions'] as Array<any>).map(RoutingInstructionTypeFromJSON)),
    };
}

export function RoutingInstructionsToChangeCriteriaRequestToJSON(value?: RoutingInstructionsToChangeCriteriaRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'compRequestInfo': CompRoutingRequestTypeToJSON(value.compRequestInfo),
        'instructions': value.instructions === undefined ? undefined : ((value.instructions as Array<any>).map(RoutingInstructionTypeToJSON)),
    };
}

