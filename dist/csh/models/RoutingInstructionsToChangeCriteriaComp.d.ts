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
import type { CompAccountingType } from './CompAccountingType';
import type { CompRoutingRequestType } from './CompRoutingRequestType';
import type { PayeeInfoType } from './PayeeInfoType';
import type { RoutingInstructionType } from './RoutingInstructionType';
/**
 * Comp Accounting Routing Info
 * @export
 * @interface RoutingInstructionsToChangeCriteriaComp
 */
export interface RoutingInstructionsToChangeCriteriaComp {
    /**
     *
     * @type {CompAccountingType}
     * @memberof RoutingInstructionsToChangeCriteriaComp
     */
    compAccountingInfo?: CompAccountingType;
    /**
     *
     * @type {CompRoutingRequestType}
     * @memberof RoutingInstructionsToChangeCriteriaComp
     */
    compRequestInfo?: CompRoutingRequestType;
    /**
     *
     * @type {PayeeInfoType}
     * @memberof RoutingInstructionsToChangeCriteriaComp
     */
    payeeInfo?: PayeeInfoType;
    /**
     * Set of routing instructions associated to this routing type.
     * @type {Array<RoutingInstructionType>}
     * @memberof RoutingInstructionsToChangeCriteriaComp
     */
    instructions?: Array<RoutingInstructionType>;
    /**
     *
     * @type {number}
     * @memberof RoutingInstructionsToChangeCriteriaComp
     */
    folioWindowNo?: number;
}
/**
 * Check if a given object implements the RoutingInstructionsToChangeCriteriaComp interface.
 */
export declare function instanceOfRoutingInstructionsToChangeCriteriaComp(value: object): boolean;
export declare function RoutingInstructionsToChangeCriteriaCompFromJSON(json: any): RoutingInstructionsToChangeCriteriaComp;
export declare function RoutingInstructionsToChangeCriteriaCompFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingInstructionsToChangeCriteriaComp;
export declare function RoutingInstructionsToChangeCriteriaCompToJSON(value?: RoutingInstructionsToChangeCriteriaComp | null): any;
