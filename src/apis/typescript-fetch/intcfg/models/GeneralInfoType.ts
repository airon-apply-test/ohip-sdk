/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { SVTransactionType } from './SVTransactionType';
import {
    SVTransactionTypeFromJSON,
    SVTransactionTypeFromJSONTyped,
    SVTransactionTypeToJSON,
} from './SVTransactionType';

/**
 * 
 * @export
 * @interface GeneralInfoType
 */
export interface GeneralInfoType {
    /**
     * Create stored value during checkin of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    createSVCheckIn?: boolean;
    /**
     * Enable failover of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    enableFailover?: boolean;
    /**
     * CC vault of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    ccVaultFunction?: boolean;
    /**
     * Enable resend of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    enableResend?: boolean;
    /**
     * Send AR auth type of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    sendARAuthType?: boolean;
    /**
     * Enable DB verification of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    enableDBVerification?: boolean;
    /**
     * Allow CC cancel transactions of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    allowCCCancelTransaction?: boolean;
    /**
     * Show stored value pin of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    showStoredValuePin?: boolean;
    /**
     * 
     * @type {SVTransactionType}
     * @memberof GeneralInfoType
     */
    showSVTransaction?: SVTransactionType;
    /**
     * Stored value system of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    storedValueSystem?: boolean;
    /**
     * Send enhanced fields of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    sendEnhancedFields?: boolean;
    /**
     * Card present of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    cardPresentFlag?: boolean;
    /**
     * Handle night audit command of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    handleNightAudit?: boolean;
    /**
     * Regular transaction command of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    regularTransaction?: boolean;
    /**
     * Courtesy card handling command of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    courtesyCardHandling?: boolean;
    /**
     * Send end of day of a Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    sendEndOfDay?: boolean;
    /**
     * Logical port number 1 of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    portNumber1?: string;
    /**
     * Logical port number 2 of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    portNumber2?: string;
    /**
     * Logical port number 3 of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    portNumber3?: string;
    /**
     * Logical port number 4 of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    portNumber4?: string;
    /**
     * IP Address of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    iPAddress?: string;
    /**
     * Stored value reedeem transaction of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    sVRedeemTrx?: string;
    /**
     * Device of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    device?: string;
    /**
     * Department code 1 of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    departmentCode1?: string;
    /**
     * Department code 2 of a the Hotel Interface
     * @type {string}
     * @memberof GeneralInfoType
     */
    departmentCode2?: string;
    /**
     * Rollup transactions flag of the Hotel Interface
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    rollupTransactions?: boolean;
    /**
     * Indicator if video checkout is allowed or not.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    allowVideoCheckout?: boolean;
    /**
     * One of the video checkout options whether to check the credit limit.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    checkCreditLimit?: boolean;
    /**
     * One of the video checkout options whether to check the balance.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    checkBalance?: boolean;
    /**
     * One of the video checkout options whether to check the routing.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    checkRouting?: boolean;
    /**
     * Indicator if the Hotel Interface handles wakeup calls.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    wakeupYn?: boolean;
    /**
     * Determines who handles the wakeup calls, either by Opera (0) or a Vendor (1).
     * @type {string}
     * @memberof GeneralInfoType
     */
    wakeupHandledBy?: string;
    /**
     * Wakeup retry limit of the Hotel Interface. Value range is 0 to 999.
     * @type {number}
     * @memberof GeneralInfoType
     */
    wakeRetry?: number;
    /**
     * Wakeup delay value of the Hotel Interface.
     * @type {number}
     * @memberof GeneralInfoType
     */
    wakeDelay?: number;
    /**
     * Used by the Interface module (IFC7.exe).
     * @type {number}
     * @memberof GeneralInfoType
     */
    prevWake?: number;
    /**
     * Folio display option of the Hotel Interface applicable to VID and MSC. "0" means folio display option off. "A" means a guest folio display option. "U" means user defined folio display option.
     * @type {string}
     * @memberof GeneralInfoType
     */
    folioDisplayOption?: string;
    /**
     * Indicator if the User defined folio 1 is utilized.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    userDefinedFolio1?: boolean;
    /**
     * Indicator if the User defined folio 2 is utilized.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    userDefinedFolio2?: boolean;
    /**
     * Indicator if the User defined folio 3 is utilized.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    userDefinedFolio3?: boolean;
    /**
     * Indicator if the User defined folio 4 is utilized.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    userDefinedFolio4?: boolean;
    /**
     * Indicator if the phone name will be displayed.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    displayPhoneName?: boolean;
    /**
     * Indicator if the Key Pin is active or not.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    keyPinActive?: boolean;
    /**
     * Indicator for Default All Print Details On.
     * @type {boolean}
     * @memberof GeneralInfoType
     */
    allPrintDetailOn?: boolean;
}

/**
 * Check if a given object implements the GeneralInfoType interface.
 */
export function instanceOfGeneralInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GeneralInfoTypeFromJSON(json: any): GeneralInfoType {
    return GeneralInfoTypeFromJSONTyped(json, false);
}

export function GeneralInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): GeneralInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'createSVCheckIn': !exists(json, 'createSVCheckIn') ? undefined : json['createSVCheckIn'],
        'enableFailover': !exists(json, 'enableFailover') ? undefined : json['enableFailover'],
        'ccVaultFunction': !exists(json, 'ccVaultFunction') ? undefined : json['ccVaultFunction'],
        'enableResend': !exists(json, 'enableResend') ? undefined : json['enableResend'],
        'sendARAuthType': !exists(json, 'sendARAuthType') ? undefined : json['sendARAuthType'],
        'enableDBVerification': !exists(json, 'enableDBVerification') ? undefined : json['enableDBVerification'],
        'allowCCCancelTransaction': !exists(json, 'allowCCCancelTransaction') ? undefined : json['allowCCCancelTransaction'],
        'showStoredValuePin': !exists(json, 'showStoredValuePin') ? undefined : json['showStoredValuePin'],
        'showSVTransaction': !exists(json, 'showSVTransaction') ? undefined : SVTransactionTypeFromJSON(json['showSVTransaction']),
        'storedValueSystem': !exists(json, 'storedValueSystem') ? undefined : json['storedValueSystem'],
        'sendEnhancedFields': !exists(json, 'sendEnhancedFields') ? undefined : json['sendEnhancedFields'],
        'cardPresentFlag': !exists(json, 'cardPresentFlag') ? undefined : json['cardPresentFlag'],
        'handleNightAudit': !exists(json, 'handleNightAudit') ? undefined : json['handleNightAudit'],
        'regularTransaction': !exists(json, 'regularTransaction') ? undefined : json['regularTransaction'],
        'courtesyCardHandling': !exists(json, 'courtesyCardHandling') ? undefined : json['courtesyCardHandling'],
        'sendEndOfDay': !exists(json, 'sendEndOfDay') ? undefined : json['sendEndOfDay'],
        'portNumber1': !exists(json, 'portNumber1') ? undefined : json['portNumber1'],
        'portNumber2': !exists(json, 'portNumber2') ? undefined : json['portNumber2'],
        'portNumber3': !exists(json, 'portNumber3') ? undefined : json['portNumber3'],
        'portNumber4': !exists(json, 'portNumber4') ? undefined : json['portNumber4'],
        'iPAddress': !exists(json, 'iPAddress') ? undefined : json['iPAddress'],
        'sVRedeemTrx': !exists(json, 'sVRedeemTrx') ? undefined : json['sVRedeemTrx'],
        'device': !exists(json, 'device') ? undefined : json['device'],
        'departmentCode1': !exists(json, 'departmentCode1') ? undefined : json['departmentCode1'],
        'departmentCode2': !exists(json, 'departmentCode2') ? undefined : json['departmentCode2'],
        'rollupTransactions': !exists(json, 'rollupTransactions') ? undefined : json['rollupTransactions'],
        'allowVideoCheckout': !exists(json, 'allowVideoCheckout') ? undefined : json['allowVideoCheckout'],
        'checkCreditLimit': !exists(json, 'checkCreditLimit') ? undefined : json['checkCreditLimit'],
        'checkBalance': !exists(json, 'checkBalance') ? undefined : json['checkBalance'],
        'checkRouting': !exists(json, 'checkRouting') ? undefined : json['checkRouting'],
        'wakeupYn': !exists(json, 'wakeupYn') ? undefined : json['wakeupYn'],
        'wakeupHandledBy': !exists(json, 'wakeupHandledBy') ? undefined : json['wakeupHandledBy'],
        'wakeRetry': !exists(json, 'wakeRetry') ? undefined : json['wakeRetry'],
        'wakeDelay': !exists(json, 'wakeDelay') ? undefined : json['wakeDelay'],
        'prevWake': !exists(json, 'prevWake') ? undefined : json['prevWake'],
        'folioDisplayOption': !exists(json, 'folioDisplayOption') ? undefined : json['folioDisplayOption'],
        'userDefinedFolio1': !exists(json, 'userDefinedFolio1') ? undefined : json['userDefinedFolio1'],
        'userDefinedFolio2': !exists(json, 'userDefinedFolio2') ? undefined : json['userDefinedFolio2'],
        'userDefinedFolio3': !exists(json, 'userDefinedFolio3') ? undefined : json['userDefinedFolio3'],
        'userDefinedFolio4': !exists(json, 'userDefinedFolio4') ? undefined : json['userDefinedFolio4'],
        'displayPhoneName': !exists(json, 'displayPhoneName') ? undefined : json['displayPhoneName'],
        'keyPinActive': !exists(json, 'keyPinActive') ? undefined : json['keyPinActive'],
        'allPrintDetailOn': !exists(json, 'allPrintDetailOn') ? undefined : json['allPrintDetailOn'],
    };
}

export function GeneralInfoTypeToJSON(value?: GeneralInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'createSVCheckIn': value.createSVCheckIn,
        'enableFailover': value.enableFailover,
        'ccVaultFunction': value.ccVaultFunction,
        'enableResend': value.enableResend,
        'sendARAuthType': value.sendARAuthType,
        'enableDBVerification': value.enableDBVerification,
        'allowCCCancelTransaction': value.allowCCCancelTransaction,
        'showStoredValuePin': value.showStoredValuePin,
        'showSVTransaction': SVTransactionTypeToJSON(value.showSVTransaction),
        'storedValueSystem': value.storedValueSystem,
        'sendEnhancedFields': value.sendEnhancedFields,
        'cardPresentFlag': value.cardPresentFlag,
        'handleNightAudit': value.handleNightAudit,
        'regularTransaction': value.regularTransaction,
        'courtesyCardHandling': value.courtesyCardHandling,
        'sendEndOfDay': value.sendEndOfDay,
        'portNumber1': value.portNumber1,
        'portNumber2': value.portNumber2,
        'portNumber3': value.portNumber3,
        'portNumber4': value.portNumber4,
        'iPAddress': value.iPAddress,
        'sVRedeemTrx': value.sVRedeemTrx,
        'device': value.device,
        'departmentCode1': value.departmentCode1,
        'departmentCode2': value.departmentCode2,
        'rollupTransactions': value.rollupTransactions,
        'allowVideoCheckout': value.allowVideoCheckout,
        'checkCreditLimit': value.checkCreditLimit,
        'checkBalance': value.checkBalance,
        'checkRouting': value.checkRouting,
        'wakeupYn': value.wakeupYn,
        'wakeupHandledBy': value.wakeupHandledBy,
        'wakeRetry': value.wakeRetry,
        'wakeDelay': value.wakeDelay,
        'prevWake': value.prevWake,
        'folioDisplayOption': value.folioDisplayOption,
        'userDefinedFolio1': value.userDefinedFolio1,
        'userDefinedFolio2': value.userDefinedFolio2,
        'userDefinedFolio3': value.userDefinedFolio3,
        'userDefinedFolio4': value.userDefinedFolio4,
        'displayPhoneName': value.displayPhoneName,
        'keyPinActive': value.keyPinActive,
        'allPrintDetailOn': value.allPrintDetailOn,
    };
}

